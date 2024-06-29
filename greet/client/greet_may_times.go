package main

import (
	"context"
	"grpc/greet/proto"
	"io"
	"log"
)

// doGreetManyTimes client for server streaming
func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := proto.GreetRequest{
		FirstName: "Navneet",
	}

	stream, err := c.GreetManyTimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break

		}

		if err!=nil{
			log.Fatalf("Error while receiving stream: %v\n", err)
		}

		log.Printf("GreetingManyTimes: %s\n",msg.Result)
	}

}
