package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

//function to call the rpc endpoint
func doGreetManyTimes(c pb.GreetServiceClient){
	log.Println("doGreet is invoked")

	req := &pb.GreetRequest{
		FirstName: "Raunak",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err !=nil{
		log.Fatalf("Error while calling the greetmany times: %v\n",err)

	}
	for {
		msg, err :=stream.Recv()
     
		if err == io.EOF {
          break
		}
		if err !=nil {
			log.Fatalf("Error while streaming : %v\n",err)
		}
		log.Printf("Greet mainy times: %v\n",msg.Result )
	}
}
