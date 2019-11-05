package main

import (
	"context"
	"github.com/Ankr-network/dccn-common/access/example/hello/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	rsp, err := c.SayHello(mockContextWithToken(), &pb.Req{Name: "ankr"})
	if err != nil {
		log.Fatalf("api error: %v", err)
	}
	log.Printf("response:%v", rsp.Message)
}

func mockContextWithToken() context.Context {
	// should be same as dccn-facade add token, not implement yet
	return context.Background()
}
