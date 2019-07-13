package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	team "github.com/Ankr-network/dccn-common/protos/teammgr/v1/grpc"
	jwt_token "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Token struct {
	Exp int64
	Jti string
	Iss string
}

var TEAM_SERVICE = "team:50051"



func GetUserID(ctx context.Context) string {
	meta, ok := metadata.FromIncomingContext(ctx)
	log.Printf("meta %+v \n", meta)
	// Note this is now uppercase (not entirely sure why this is...)
	var token string
	if ok {
		tokenArray := meta["authorization"]
		log.Printf("GetUserID: Authorization is : ", tokenArray)
		if tokenArray != nil {
			token = tokenArray[0]
		} else {
			return ""
		}
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

	return dat.Jti
}

func GetUserIDAndTeamID(ctx context.Context)(string, string) {
	meta, ok := metadata.FromIncomingContext(ctx)
	log.Printf("meta %+v \n", meta)
	// Note this is now uppercase (not entirely sure why this is...)
	var token string
	if ok {
		tokenArray := meta["authorization"]
		log.Printf("GetUserID: Authorization is : ", tokenArray)
		if tokenArray != nil {
			token = tokenArray[0]
		} else {
			return "", ""
		}
	}
	log.Printf("token %+v \n", token)

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		log.Printf("GetUserID parse error. format error token %s => %+v \n", token, parts)
		return "", ""
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

	uid := dat.Jti
	team_id := getTeamID(uid)
	return uid, team_id
}

func getTeamID(uid string) string {
	conn, err := grpc.Dial(TEAM_SERVICE, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Println(err.Error())
		}
	}(conn)

	teamClient := team.NewTeamMgrClient(conn)

	//task.Attributes.R = 1
	//t := common_proto.APP_TypeDeployment{TypeDeployment: &common_proto.TaskTypeDeployment{Image:"nginx:1.12"}}
	//task.TypeData = &t

	if rsp, err := teamClient.GetUserTeamID(context.Background(), &team.UserID{Uid: uid}); err != nil {
		//log.Println("detail create %+v " + rsp)
		log.Printf("GetUserTeamID  failed, error : %s \n", err)
		return ""
	} else {
		log.Printf("GetUserTeamID  successfully team_id : %s  \n", rsp.TeamId)
		return rsp.TeamId
	}
}
