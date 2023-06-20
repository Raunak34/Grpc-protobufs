package main

import (
	"context"
	"log"

	

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNo: 1,
		SecondNo: 1,
	})
	if err!=nil {
		log.Fatalf("could not sum: %v",err)
	}
	log.Printf("sum is: %d\n",res.Result)
}
