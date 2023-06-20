package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--read blog was invoked--")

	req := &pb.BlogId{Id: id}
    res, err  := c.ReadBlog(context.Background(), req)

	if err != nil {
		log. Printf("error happened while reading:%v\n",err)
	}
	log. Printf(" blog was read: %v\n",res)
	return res
}