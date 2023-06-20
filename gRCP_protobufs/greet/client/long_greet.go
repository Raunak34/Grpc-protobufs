package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)
func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet was invoked")
    reqs := []*pb.GreetRequest {  //array of greetrequest  
         {FirstName: "Raunak"},
		 {FirstName: "Marie"},
		 {FirstName: "Gaurav"},

	}
	stream, err := c.LongGreet(context.Background())
	if err!=nil {
		log.Fatalf("error while calling the long greet %v\n",err)
	}

	for _, req := range reqs{
		log.Printf("sending req : %v\n",req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err :=stream.CloseAndRecv()
        if err!=nil {
			log.Fatalf("error while receiving reponse fom long greet: %v\n",err)
		}
		log.Printf("longGreet :%s\n",res.Result)
	}