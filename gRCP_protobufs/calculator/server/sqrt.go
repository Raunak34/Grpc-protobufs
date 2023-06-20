package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
func (s * Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error){
   log.Printf("sqrt func was invoked: %v\n",in)
   number := in.Number

   if number < 0 {
	return nil, status.Errorf(
		codes.InvalidArgument,
		fmt.Sprintf("received a negetive number, %d\n",number),
	)
   }
   return &pb.SqrtResponse{
	Result: math.Sqrt(float64(number)),
   },nil
}