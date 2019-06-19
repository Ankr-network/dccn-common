package api

import (
	"context"
	"log"
)

// Hello server
// Server represents the gRPC server
type Server struct {
}

// Ping generates response to a Ping request
func (s *Server) Ping(ctx context.Context, in *PingMsg) (*PingMsg, error) {
	log.Printf("Receive message: %s", in.Id)
	return &PingMsg{Id: "Hello " + in.Id}, nil
}
