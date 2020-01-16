package main

import (
	"context"
	"fmt"
	"github.com/Ankr-network/dccn-common/auth"
	"github.com/Ankr-network/dccn-common/auth/example/hello/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct{}

func (p *service) SayHello(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	log.Printf("SayHello receive: %v", req.Name)

	claim, err := auth.GetClaim(ctx)
	if err != nil {
		log.Printf("GetClaim error: %v", err)
	}
	log.Printf("claim: %+v", claim)

	uid, err := auth.GetUID(ctx)
	if err != nil {
		log.Printf("GetUID error: %v", err)
	}
	log.Printf("uid: %+v", uid)

	cid, err := auth.GetClientID(ctx)
	if err != nil {
		log.Printf("GetClientID error: %v", err)
	}
	log.Printf("cid: %+v", cid)

	return &pb.Rsp{
		Message: fmt.Sprintf("hello %s", req.Name),
	}, nil
}

func (p *service) SayHelloInsecure(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	log.Printf("SayHelloInsecure receive: %v", req.Name)

	return &pb.Rsp{
		Message: fmt.Sprintf("insecure hello %s", req.Name),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	verifier, err := auth.NewVerifier(auth.ExcludeMethods("/pb.Hello/SayHelloInsecure"))
	if err != nil {
		log.Fatalf("newVerifier error:%v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(verifier.GRPCUnaryInterceptor()))
	pb.RegisterHelloServer(s, &service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
