package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)
func doGreetEveryone(c pb.GreetServiceClient) {
 log.Println("dogreet function was invoked")
 stream, err :=c.GreetEveryone(context.Background()) 

 if err!=nil{
	log.Fatalf("err while creating stream %v\n",err)

 }
 reqs := []*pb.GreetRequest{
	{FirstName: "Raunak"},
	{FirstName: "Soumik"},
	{FirstName: "Rubai"},
 }
 waitc :=make(chan struct{})
 go func ()  {
	for _, req :=range reqs {
		log.Printf("send request : %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
 }() //() this is to invoke
 go func() {
	for {
		res, err :=stream.Recv()

		if err==io.EOF{
			break
		}
		if err!=nil {
			log.Printf("Error while receiving: %v\n",err)
			break
		}
        log.Printf("Received  %v\n",res.Result)
	}
	close(waitc)
 }()
<-waitc //this will wait for the two go routines to exchange all the request and rsponse


}