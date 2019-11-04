package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/dccn-common/protos/usermgr/v1/grpc"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Token struct {
	Exp int64
	Jti string
	Iss string
}

const (
	userMgrEndpoint = "usermgr:50051"
)

func GetUserID(ctx context.Context) string {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("error: cannot get metadata from ctx")
		return ""
	}

	tokenStringArray := meta["authorization"]
	if tokenStringArray == nil || len(tokenStringArray) == 0 {
		log.Printf("error: autorization header is missing")
		return ""
	}

	tokenString := tokenStringArray[0]

	tokenParts := strings.Split(tokenString, ".")

	if len(tokenParts) != 3 {
		log.Printf("error: parse token format error %s => %+v", tokenString, tokenParts)
		return ""
	}

	raw, err := jwt.DecodeSegment(tokenParts[1])
	if err != nil {
		log.Printf("jwt.DecodeSegment error: %v", err)
	}

	var token Token

	if err := json.Unmarshal(raw, &token); err != nil {
		fmt.Println("Unmarshal error:", err)
	}

	return token.Jti
}

func GetUserIDAndTeamID(ctx context.Context) (string, string) {
	uid := GetUserID(ctx)
	teamID := GetTeamID(ctx, uid)
	return uid, teamID
}

func GetTeamID(ctx context.Context, uid string) string {
	conn, err := grpc.Dial(userMgrEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Printf("grpc.Dial %s error: %v", userMgrEndpoint, err)
	}

	client := usermgr.NewUserMgrClient(conn)

	rsp, err := client.GetUser(ctx, &usermgr.GetUserRequest{Uid: uid})
	if err != nil {
		log.Printf("grpc GetUser error: %v", err)
		return ""
	}
	return rsp.CurrentTeamId
}
