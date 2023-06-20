package main

import (
	"log"
	

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

var addr string ="localhost:50051"

func main() {
	tls := true //it can be change according to the use 
	opts := []grpc.DialOption{}
	if tls {
		certFile := "ssl/ca.crt"
		
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("error while loading the CA certificate: %v\n",err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("connnection problem: %v\n",err)  //the defer keyword is used to delay the execution of a function or a statement until the nearby function returns
	}
	defer conn.Close()
    c :=pb.NewGreetServiceClient(conn)
	doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	//doGreetEveryone(c)
	//doGreetWithDeadline(c, 1*time.Second)
}