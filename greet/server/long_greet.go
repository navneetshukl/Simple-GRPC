package main

import (
	"fmt"
	"grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res:=""

	for{
		req,err:=stream.Recv()

		if err==io.EOF{
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}
		if err!=nil{
            log.Fatalf("Error while receiving data from client: %v",err)
            return err
        }
		res+=fmt.Sprintf("Hello %s!\n",req.FirstName)
	}
}
