package main 

import (
	"context"
	"log"
	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"

)
func (s *Server)  Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {  //here the context named ctx and the requst named in
log.Printf("greet function was invoked with %v\n", in)  //in is the request 
  return &pb.GreetResponse{
	Result: "Hello " + in.FirstName,
  },nil
}
  //in is the request