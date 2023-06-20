package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)
var collection *mongo.Collection
var addr string = "0.0.0.0:50051" //localhost

type Server struct{ //we will use this server to use all the rpc endpoints that we will define in our greet proto

	pb.BlogServiceServer
}

func main () {    
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err !=nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err!=nil {
		log.Fatal(err)
	}
	collection = client.Database("blogdb").Collection("blog")
	lis, err :=  net.Listen("tcp",addr)

	if err != nil {
		log.Fatalf("failed to listen on %v\n",err)
	}
	log.Printf("listening on %s\n",addr)

	s:= grpc.NewServer()
	//register the new greet server
	pb.RegisterBlogServiceServer(s, &Server{})

	if err := s.Serve(lis); err!=nil {
		log.Fatalf("failed to serve : %v\n",err)
	}
}