package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)
func updateBlog(c pb.BlogServiceClient, id string ){
	log.Println("Update blog was invoked")

	newBlog := &pb.Blog{
		Id: id,
		AuthorId: "Not Raunak",
		Title: "A new Title",
		Content: "content of the first blog with some additions",

	}
	_, err:= c.UpdateBlog(context.Background(), newBlog)


	if err != nil {
		log.Fatalf("Error happened while updateding %v\n",err)
	}
	log.Printf("Bog was updated")
}
