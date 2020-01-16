package main

import (
	"context"
	"fmt"
	"github.com/Ankr-network/dccn-common/access"
	"github.com/Ankr-network/dccn-common/access/example/hello/pb"
	"github.com/Ankr-network/dccn-common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	sayHelloResource = "ankr:hello:test-team1:hello"
	sayHelloAction   = "say"
)

type service struct{}

func (p *service) SayHello(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	log.Printf("receive: %v", req.Name)

	if err := checkSayHelloAccess(ctx); err != nil {
		return nil, err
	}
	return &pb.Rsp{
		Message: fmt.Sprintf("hello %s", req.Name),
	}, nil
}

func checkSayHelloAccess(ctx context.Context) error {
	uid := util.GetUserID(ctx)
	ok, err := access.Authorize(ctx, uid, sayHelloResource, sayHelloAction)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if !ok {
		return status.Error(codes.PermissionDenied, "access denied")
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
