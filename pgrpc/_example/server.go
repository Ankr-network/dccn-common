package main

import (
	"log"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	grpc "google.golang.org/grpc"
)

func server() {
	ln, err := pgrpc.Listen("tcp", "127.0.0.1:50051", genHook("server"))
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	api.RegisterPingServer(server, &api.Server{})
	if err := server.Serve(ln); err != nil {
		log.Fatalln(err)
	}
}
