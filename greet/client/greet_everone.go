package main

import (
	"context"
	"grpc/greet/proto"
	"io"
	"log"
	"time"
)

// doGreetEveryOne client code for bidirectional streaming
func doGreetEveryOne(c proto.GreetServiceClient) {
	log.Println("doGreetEveryOne was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Alice"},
		{FirstName: "Bob"},
		{FirstName: "Charlie"},
		{FirstName: "David"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)

	}()

	<-waitc

}
