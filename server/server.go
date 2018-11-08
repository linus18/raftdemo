package server

import (
	"context"
	"fmt"

	raft "github.com/linus18/raft/pb"
)

// Transport encapsulates RPC channel facilitating communication between nodes
type Transport struct {
	RPC chan interface{}
}

// NewTransport creates a new Transport struct
func NewTransport(rpc chan interface{}) *Transport {
	return &Transport{
		RPC: rpc,
	}
}

// RequestVote RPC handler
func (t *Transport) RequestVote(ctx context.Context, in *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	fmt.Printf("(term=%d, candidate=%v)\n", in.Term, string(in.Candidate))
	t.RPC <- in
	out := <-t.RPC
	resp, _ := out.(raft.RequestVoteResponse)
	return &resp, nil
}

// AppendEntries RPC handler
func (t *Transport) AppendEntries(tx context.Context, in *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error) {
	t.RPC <- in
	out := <-t.RPC
	resp, _ := out.(raft.AppendEntriesResponse)
	return &resp, nil
}
