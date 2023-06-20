package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)
func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error{
	log.Println("avg function was invoked")
	var sum int32 = 0
	count:=0
	 for {
		req, err :=stream.Recv()//returns a lenght of message or daatgram in bytes
		if err==io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
            Result: float64(sum)/float64(count),
			})
		}
		if err!=nil {
			log.Fatalf("error while reading client stream: %v\n",err)
		}
		log.Printf("receiving number :%d\n", req.Number)
		sum += req.Number
		count++
	}
	 }
