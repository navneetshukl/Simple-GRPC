package main

import (
	"context"
	"grpc/greet/proto"
	"log"
	"time"
)

// doLongGreet function for client side streaming
func doLongGreet(c proto.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*proto.GreetRequest{
		{FirstName: "Navneet"},
		{FirstName: "Yatinjal"},
		{FirstName: "Chotu"},
		{FirstName: "Don"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)

		time.Sleep(1 * time.Second)

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet Result: %s\n", res.Result)

}
