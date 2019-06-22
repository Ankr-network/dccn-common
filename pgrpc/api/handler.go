package api

import (
	"context"
)

// Hello server
// Server represents the gRPC server
type Server struct{}

// Ping generates response to a Ping request
func (s *Server) Ping(ctx context.Context, in *PingMsg) (*PingMsg, error) {
	return &PingMsg{Id: "Hello " + in.Id}, nil
}
