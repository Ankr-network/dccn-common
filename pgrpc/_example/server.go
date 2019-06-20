package main

import (
	"log"
	"net"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/api"
	grpc "google.golang.org/grpc"
)

func server() {
	ln, err := pgrpc.Listen("tcp", "127.0.0.1:50051", "example_server", func(conn *net.Conn, err error) {
		if err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	api.RegisterPingServer(server, &api.Server{})
	if err := server.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}
