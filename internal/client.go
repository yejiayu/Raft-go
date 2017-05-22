package internal

import (
	"log"

	"google.golang.org/grpc"
)

func newClient(host string) RaftClient {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	log.Println("new RPC client", host)

	return NewRaftClient(conn)
}
