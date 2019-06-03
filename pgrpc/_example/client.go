package main

import (
	context "context"
	"log"
	"strconv"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	grpc "google.golang.org/grpc"
)

func client() {
	if err := pgrpc.InitClient("tcp", ":50051", nil, grpc.WithInsecure()); err != nil {
		log.Fatalln(err)
	}

	// wait server connect
	time.Sleep(1 << 30)

	var oneKey string
	{ // test loop all
		pgrpc.Each(func(key string, conn *grpc.ClientConn) error {
			resp, err := api.NewPingClient(conn).SayHello(context.Background(), &api.PingMessage{
				Greeting: "Hello " + key,
			})
			if err != nil {
				log.Fatalln(err)
				return err
			}

			log.Println(resp.Greeting)

			oneKey = key
			return nil
		})
	}

	{ // test dial
		cc, err := pgrpc.Dial(oneKey)
		if err != nil {
			log.Fatalln(err)
		}
		resp, err := api.NewPingClient(cc).SayHello(context.Background(), &api.PingMessage{
			Greeting: "dial",
		})
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resp.Greeting)
		cc.Close()
	}
	{ // test alias
		pgrpc.Alias(oneKey, "test")
		cc, err := pgrpc.Dial("test")
		if err != nil {
			log.Fatalln(err)
		}
		defer cc.Close()

		for i := 0; i < 2; i++ {
			resp, err := api.NewPingClient(cc).SayHello(context.Background(), &api.PingMessage{
				Greeting: "test-" + strconv.Itoa(i),
			})
			if err != nil {
				log.Fatalln(err)
			}

			log.Println(resp.Greeting)
		}
	}
}
