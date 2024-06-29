package main

import (
	"context"
	"grpc/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient){
	log.Println("doGreet was invoked")

	req:=proto.GreetRequest{
		FirstName: "Navneet",
	}

	res,err:=c.Greet(context.Background(),&req)

	if err!=nil{
		log.Fatalf("Error while calling Greet: %v",err)
	}

	log.Printf("Greeting: %s\n",res.Result)
}