package internal

import (
	"log"
	"testing"

	context "golang.org/x/net/context"
)

const (
	port = ":50000"
	host = "localhost:50000"
)

func init() {
	go newServer(port)
}
func TestRPC(t *testing.T) {
	c := newClient(host)

	appendRes, err := c.AppendEntries(context.Background(), &AppendRequest{Name: "append"})
	if err != nil || appendRes.Message != "append" {
		t.Fail()
	}

	voteRes, err := c.RequestVote(context.Background(), &VoteRequest{Name: "vote"})
	if err != nil || voteRes.Message != "vote" {
		t.Fail()
	}

	log.Println("done")
}
