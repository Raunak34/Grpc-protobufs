package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

    pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

var addr string ="localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("connnection problem: %v\n",err)  //the defer keyword is used to delay the execution of a function or a statement until the nearby function returns
	}
	defer conn.Close()
    c :=pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	//doAvg(c)
	//doMax(c)
	doSqrt(c, -10)
}