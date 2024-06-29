package main

import (
	"grpc/greet/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	log.Println("Connected to grpc server")

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	//doGreet(client)  // code for unary client
	//doGreetManyTimes(client) // code for server streaming client
	doLongGreet(client) //code for client side streaming

}
