package internal

import (
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) RequestVote(ctx context.Context, req *VoteRequest) (*VoteResponse, error) {
	return &VoteResponse{Message: req.Name}, nil
}

func (s *server) AppendEntries(ctx context.Context, req *AppendRequest) (*AppendResponse, error) {
	return &AppendResponse{Message: req.Name}, nil
}

func newServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	defer lis.Close()
	s := grpc.NewServer()

	RegisterRaftServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
