package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)
func doMax (c pb.CalculatorServiceClient){
	log.Println("Domax function was invoked")
	stream, err:=c.Max(context.Background())
	 if err !=nil {
		log.Fatalf("Error while opening the stream: %vv\n",err)
	 }
	 waitc  := make (chan struct{})
	  go func() {
         numbers:= [] int32{4,7,2,19,4,6,32}
		 for _, number := range numbers {
            log.Printf("sending the numbers: %d\n",number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})
			time.Sleep(1 * time.Second)
		 }
		 stream.CloseSend()
	  } ()
	  go func() {
		for{
			res, err :=stream.Recv()
			if err == io.EOF {
				break
			}
			if err !=nil {
				log.Printf("problem while reading the server stream %v\n",err)
				break
			}
			log.Printf("Rceived a new maximum: %d\n",res.Result)
		}
		close(waitc)
	  }()
      <-waitc
}