package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jonnypillar/grpc-example/api"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	resposne, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "Foo"})
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	fmt.Println("Response from server: ", resposne)
}
