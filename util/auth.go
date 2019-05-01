package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"strings"
	jwt_token "github.com/dgrijalva/jwt-go"
)


type Token struct {
	Exp int64
	Jti string
	Iss string
}


func GetUserID(ctx context.Context) string{
	meta, ok := metadata.FromContext(ctx)
	// Note this is now uppercase (not entirely sure why this is...)
	var token string
	if ok {
		token = meta["token"]
	}
	parts := strings.Split(token, ".")

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
