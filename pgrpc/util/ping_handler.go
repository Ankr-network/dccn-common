package util

import (
	"context"
	"time"
	grpc "google.golang.org/grpc"
)

func PingHook(cc *grpc.ClientConn) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := NewPingClient(cc).Ping(ctx, &PingMsg{Id: "ping"})
	cancel()
	return err
}

// Server represents the gRPC server
type Server struct{}

// Ping generates response to a Ping request
func (s *Server) Ping(ctx context.Context, in *PingMsg) (*PingMsg, error) {
	return &PingMsg{Id: "Hello " + in.Id}, nil
}
