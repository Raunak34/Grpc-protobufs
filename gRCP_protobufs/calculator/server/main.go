package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051" //localhost

type Server struct{ //we will use this server to use all the rpc endpoints that we will define in our greet proto

	pb.CalculatorServiceServer
}

func main () {    
	lis, err :=  net.Listen("tcp",addr)

	if err != nil {
		log.Fatalf("failed to listen on %v\n",err)
	}
	log.Printf("listening on %s\n",addr)

	s:= grpc.NewServer()
	//register the new greet server
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err!=nil {
		log.Fatalf("failed to serve : %v\n",err)
	}
}