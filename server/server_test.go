package server

import "testing"

func Test(t *testing.T) {
	transport := NewTransport(make(chan interface{}))
	if transport == nil {
		t.Error("failed")
	}
}
