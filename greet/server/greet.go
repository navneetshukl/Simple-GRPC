package main

import (
	"context"
	"grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v\n", req)

	return &proto.GreetResponse{
		Result: "Hello" + req.FirstName,
	}, nil

}
