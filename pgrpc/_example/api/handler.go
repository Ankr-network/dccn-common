package api

import (
	"context"
	"log"
)

// Hello server
// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message: %s", in.Greeting)
	return &PingMessage{Greeting: "Hello " + in.Greeting}, nil
}
