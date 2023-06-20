package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient , n int32){
	log.Println("Do sqrt was inkoed")
	res,err :=c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n})
		if err !=nil {
			e, ok := status.FromError(err)
			if ok {
				log.Printf("Error message from the server : %s\n",e.Message())
				log.Printf("rror code from the server: %s\n",e.Code())

				if e.Code() == codes.InvalidArgument {
					log.Println("we have a negetive number")
					return
				}
			} else {
				log.Fatalf("A non grpc error %v\n",err)
			}
			
		}
		log.Printf("Sqrt of %f\n",res.Result)
}