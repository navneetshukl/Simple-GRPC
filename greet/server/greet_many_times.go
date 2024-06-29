package main

import (
	"fmt"
	"grpc/greet/proto"
	"log"
)

// GreetManyTimes server side streaming
func (s *Server) GreetManyTimes(req *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("Greet many times function was invoked with : %v\n", req)

	for i := 0; i < 10; i++ {

		res := fmt.Sprintf("Hello %s, number %d ", req.FirstName, i)
		err := stream.Send(&proto.GreetResponse{Result: res})
		if err != nil {
			log.Println("Error in sending the response ", err)
			return err
		}

	}
	return nil
}
