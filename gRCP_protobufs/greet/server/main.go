package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:50051" //localhost

type Server struct{ //we will use this server to use all the rpc endpoints that we will define in our greet proto

	pb.GreetServiceServer
}

func main () {    
	lis, err :=  net.Listen("tcp",addr)

	if err != nil {
		log.Fatalf("failed to listen on %v\n",err)
	}
	log.Printf("listening on %s\n",addr)
    opts := []grpc.ServerOption{}
	tls := true   // you can change that according to your use 
	 if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failedl loading certificates: %v\n",err)
		}
		opts = append(opts, grpc.Creds(creds))
	 }

	s:= grpc.NewServer(opts...)
	//register the new greet server
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err!=nil {
		log.Fatalf("failed to serve : %v\n",err)
	}
}