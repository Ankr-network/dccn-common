package main

import (
	context "context"
	"log"
	"net"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/util"
	grpc "google.golang.org/grpc"
)

func client() {
	if err := pgrpc.InitClient("tcp", ":50051", func(conn *net.Conn) string {
		var key string
		var err error
		*conn, key, err = util.ParseProxyProto(*conn)
		if err != nil {
			log.Fatalln(err)
		}
		return key
	}, util.PingHook, grpc.WithInsecure()); err != nil {
		log.Fatalln(err)
	}

	// wait server connect
	time.Sleep(1 << 30)

	var oneKey string
	{ // test loop all
		pgrpc.Each(func(key string, ips []string, conn *grpc.ClientConn, err error) error {
			if err != nil {
				log.Println(err)
				return err
			}

			resp, err := util.NewPingClient(conn).Ping(context.Background(), &util.PingMsg{
				Id: "Hello " + key,
			})
			if err != nil {
				log.Fatalf("%+v", err)
				return err
			}

			log.Println(resp.Id, ips)

			oneKey = key
			return nil
		})
	}

	if oneKey == "" {
		return
	}
	// test dial
	cc, err := pgrpc.Dial(oneKey)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := util.NewPingClient(cc).Ping(context.Background(), &util.PingMsg{
		Id: "dial",
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Id)
	cc.Close()
}
