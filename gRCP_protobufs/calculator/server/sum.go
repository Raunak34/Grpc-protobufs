package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error){
	log.Printf("sum function was invoked %v\n",in)
	return &pb.SumResponse{
		Result: in.FirstNo + in.SecondNo,
	},nil
}