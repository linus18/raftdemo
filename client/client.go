package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	raft "github.com/linus18/raft/pb"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("address", "localhost:3001", "server address")
)

func randomTimeout(minVal time.Duration) <-chan time.Time {
	if minVal == 0 {
		return nil
	}
	extra := (time.Duration(rand.Int63()) % minVal)
	return time.After(minVal + extra)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not get a dial tone")
	}
	defer conn.Close()
	client := raft.NewLeaderElectionClient(conn)
	request := &raft.RequestVoteRequest{
		Term:      1,
		Candidate: []byte("node1"),
	}

	timer := randomTimeout(1000 * time.Millisecond)
	for {
		select {
		case <-timer:
			timer = randomTimeout(1000 * time.Millisecond)
			resp, err := client.RequestVote(context.Background(), request)
			if err != nil {
				log.Print("call failed")
			}
			fmt.Println(resp)
		}
	}
}
