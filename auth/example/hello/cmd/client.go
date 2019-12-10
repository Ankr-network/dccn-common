package main

import (
	"context"
	"github.com/Ankr-network/dccn-common/auth/example/hello/pb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	conf := oauth2.Config{
		Endpoint: oauth2.Endpoint{
			TokenURL: "http://localhost:50053/token",
		},
		Scopes: []string{"all"},
	}
	token, err := conf.PasswordCredentialsToken(context.TODO(), "test@ankr.com", "test")
	if err != nil {
		log.Fatalf("get access token error: %v", err)
	}

	cred := newPerRPCCredentials(token)

	rsp, err := c.SayHello(context.TODO(), &pb.Req{Name: "ankr"}, grpc.PerRPCCredentials(cred))
	if err != nil {
		log.Fatalf("SayHello api error: %v", err)
	}
	log.Printf("SayHello response:%v", rsp.Message)

	rsp, err = c.SayHelloInsecure(context.TODO(), &pb.Req{Name: "ankr"})
	if err != nil {
		log.Fatalf("SayHelloInsecure api error: %v", err)
	}
	log.Printf("SayHelloInsecure response:%v", rsp.Message)
}

func newPerRPCCredentials(token *oauth2.Token) credentials.PerRPCCredentials {
	return &insecure{token: token}
}

type insecure struct {
	token *oauth2.Token
}

func (p *insecure) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": p.token.Type() + " " + p.token.AccessToken,
	}, nil
}

func (p *insecure) RequireTransportSecurity() bool {
	return false
}
