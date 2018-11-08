package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/linus18/raft/pb"
	"github.com/linus18/raft/server"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 3000, "port to listen for raft communication")
	peerAddr = flag.String("peer", "localhost:3001", "address of a peer")
	nodeID   = flag.String("id", "node1", "raft node ID")
)

// RaftState captures the state of a Raft node: Follower, Candidate, Leader,
type RaftState uint8

const (
	// Follower is the initial state of a Raft node.
	Follower RaftState = iota
	// Candidate is one of the valid states of a Raft node
	Candidate
	// Leader is one of the valid states of a Raft node
	Leader
)

// Raft holds state needed by each Raft node
type Raft struct {
	TermID       uint64
	CurrentState RaftState
	Leader       []byte
	rpcCh        chan interface{}
	voted        bool
	lastContact  time.Time
}

// NewRaft constructs a new Raft node
func NewRaft(rpc chan interface{}) *Raft {
	return &Raft{
		TermID:       1,
		CurrentState: Follower,
		Leader:       nil,
		rpcCh:        rpc,
		voted:        false,
		lastContact:  time.Time{},
	}
}

func randomTimeout(minVal time.Duration) <-chan time.Time {
	if minVal == 0 {
		return nil
	}
	extra := (time.Duration(rand.Int63()) % minVal)
	return time.After(minVal + extra)
}

func serve(t *server.Transport) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("gRPC server listening on port %d...\n", *port)
	grpcServer := grpc.NewServer()

	pb.RegisterLeaderElectionServer(grpcServer, t)
	grpcServer.Serve(lis)
}

func main() {
	flag.Parse()
	rpc := make(chan interface{})
	transport := server.NewTransport(rpc)
	go serve(transport)

	r := NewRaft(rpc)

	for {
		switch r.CurrentState {
		case Follower:
			r.runFollower()
		case Candidate:
			r.runCandidate()
		case Leader:
			r.runLeader()
		}
	}
}

func (r *Raft) runFollower() {
	printState("FOLLOWER")
	heartbeat := randomTimeout(1000 * time.Millisecond)
	for {
		select {
		case <-heartbeat:
			heartbeat = randomTimeout(1000 * time.Millisecond)
			if time.Now().Sub(r.lastContact) < 1000*time.Millisecond {
				log.Println("we were contacted by the leader recently, so continue following...")
			} else {
				log.Println("no contact from leader in a while: becomming a candidate...")
				r.CurrentState = Candidate
				return
			}
		case in := <-r.rpcCh:
			r.processRPC(in)
		}
	}
}

func (r *Raft) runCandidate() {
	printState("CANDIDATE")
	if requestVote(r.TermID, *peerAddr) {
		log.Println("got enough votes: becoming a leader!")
		r.TermID++
		r.CurrentState = Leader
		return
	}
	electionTimeout := randomTimeout(1000 * time.Millisecond)
	for {
		select {
		case <-electionTimeout:
			log.Println("timer expired...keep waiting...")
			electionTimeout = randomTimeout(1000 * time.Millisecond)
			if requestVote(r.TermID, *peerAddr) {
				log.Println("got enough votes: becoming a leader!")
				r.TermID++
				r.CurrentState = Leader
				return
			}
		case in := <-r.rpcCh:
			r.processRPC(in)
			if r.CurrentState != Candidate {
				log.Println("leader has been elected!")
				return
			}
		}
	}
}

func (r *Raft) runLeader() {
	printState("LEADER")
	heartbeat := randomTimeout(100 * time.Millisecond)
	for {
		select {
		case <-heartbeat:
			heartbeat = randomTimeout(100 * time.Millisecond)
			sendHeartbeat(r.TermID, *peerAddr)
		case in := <-r.rpcCh:
			r.processRPC(in)
		}
	}
}

func (r *Raft) processRPC(in interface{}) {
	switch cmd := in.(type) {
	case *pb.RequestVoteRequest:
		fmt.Printf("(term=%d, candidate=%v)\n", cmd.Term, string(cmd.Candidate))
		granted := !(r.voted || r.TermID > cmd.Term)
		if granted {
			r.voted = true
		}
		r.rpcCh <- pb.RequestVoteResponse{
			Term:    r.TermID,
			Granted: granted,
		}
	case *pb.AppendEntriesRequest:
		fmt.Printf("AppendEntries RPC received! (term=%d)\n", cmd.Term)
		if r.TermID > cmd.Term || r.CurrentState == Leader {
			// ignore heartbeat from old term
			r.rpcCh <- pb.AppendEntriesResponse{
				Term:    r.TermID,
				Success: false,
			}
			return
		}
		r.Leader = cmd.Leader
		r.TermID = cmd.Term
		r.CurrentState = Follower
		r.lastContact = time.Now()
		r.rpcCh <- pb.AppendEntriesResponse{
			Term:    r.TermID,
			Success: true,
		}
	default:
		fmt.Println("unknown msg received")
	}
}

func requestVote(termID uint64, addr string) bool {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not get a dial tone")
	}
	defer conn.Close()
	client := pb.NewLeaderElectionClient(conn)
	request := &pb.RequestVoteRequest{
		Term:      termID,
		Candidate: []byte(*nodeID),
	}
	resp, err := client.RequestVote(context.Background(), request)
	if err != nil {
		log.Print("call failed")
		return false
	}
	if resp != nil {
		fmt.Println(resp)
		return resp.Granted
	}
	return false
}

func sendHeartbeat(termID uint64, addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not get a dial tone")
	}
	defer conn.Close()
	client := pb.NewLeaderElectionClient(conn)
	request := &pb.AppendEntriesRequest{
		Term:   termID,
		Leader: []byte(*nodeID),
	}
	resp, err := client.AppendEntries(context.Background(), request)
	if err != nil {
		log.Println("[WARN] call failed: follower must've died...")
	}
	if resp != nil {
		if !resp.Success {
			log.Println("[WARN] hm follower refused to append...am I out of date?")
		}
		fmt.Println(resp)
	}
}

func printState(state string) {
	log.Println("*****************")
	log.Printf("I AM A %v", state)
	log.Println("*****************")
}
