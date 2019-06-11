package util

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	jwt_token "github.com/dgrijalva/jwt-go"
)


type Token struct {
	Exp int64
	Jti string
	Iss string
}


func GetUserID(ctx context.Context) string{
	meta, ok := metadata.FromIncomingContext(ctx)
	log.Printf("meta %+v \n", meta)
	// Note this is now uppercase (not entirely sure why this is...)
	var token string
	if ok {
		tokenArray := meta["token"]
		token = tokenArray[0]
	}
	log.Printf("token %+v \n", token)


	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		log.Printf("GetUserID parse error. format error token %s => %+v \n", token, parts)
		return ""
	}

	decoded, err := jwt_token.DecodeSegment(parts[1])
	if err != nil {
		fmt.Println("decode error:", err)

	}
	fmt.Println(string(decoded))
	var dat Token

	if err := json.Unmarshal(decoded, &dat); err != nil {
		fmt.Println("Unmarshal error:", err)
	}
	return string(dat.Jti)
}
