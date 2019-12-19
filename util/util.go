package util

import (
	"context"
	"log"

	"github.com/Ankr-network/dccn-common/auth"
	usermgr "github.com/Ankr-network/dccn-common/protos/usermgr/v1/grpc"
	"google.golang.org/grpc"
)

type Token struct {
	Exp int64
	Jti string
	Iss string
}

var (
	verifier auth.Verifier
)

func init() {
	v, err := auth.NewVerifier()
	if err != nil {
		log.Fatalf("auth.NewVerifier error: %v", err)
	}
	verifier = v
}

const (
	userMgrEndpoint = "usermgr:50051"
)

// Deprecated: use auth.GetUID instead
func GetUserID(ctx context.Context) string {
	ctx, err := verifier.VerifyContext(ctx)
	if err != nil {
		log.Printf("verifier.VerifyContext error: %v", err)
		return ""
	}

	uid, err := auth.GetUID(ctx)
	if err != nil {
		log.Printf("auth.GetUID error: %v", err)
		return ""
	}
	return uid
}

// Deprecated: current team id should not get from backend
func GetUserIDAndTeamID(ctx context.Context) (string, string) {
	uid := GetUserID(ctx)
	teamID := GetTeamID(ctx, uid)
	return uid, teamID
}

// Deprecated: current team id should not get from backend
func GetTeamID(ctx context.Context, uid string) string {
	conn, err := grpc.Dial(userMgrEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Printf("grpc.Dial %s error: %v", userMgrEndpoint, err)
		return ""
	}

	client := usermgr.NewUserMgrClient(conn)

	rsp, err := client.GetUser(ctx, &usermgr.GetUserRequest{Uid: uid})
	if err != nil {
		log.Printf("grpc GetUser error: %v", err)
		return ""
	}
	return rsp.CurrentTeamId
}
