package main

import (
	"log"
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)
func createBlog(c pb.BlogServiceClient) string {
	log.Println("---create blog was invoked")

	blog := &pb.Blog {
		AuthorId: "Raunak",
		Title: "My First Blog",
		Content: "Content of my first Blog",
	}
	res, err :=c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("An unexpected error:%v\n",err)
	}
	log.Printf("Blog has been created:%s\n",res.Id)
	return res.Id
}