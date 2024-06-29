package main

import (
	"grpc/greet/proto"
	"io"
	"log"
)

// GreetEveryone server code for bidirectional streaming
func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to read client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + " !"

		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Failed to send response to client: %v\n", err)
		}

	}

}
