package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"log"
	"net"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("About to handle sum request")
	sumResponse := calculatorpb.SumResponse{
		SumResult: request.FirstNumber + request.SecondNumber,
	}
	return &sumResponse, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to create server. ", err)
	}
	log.Println("created a server")
	grpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(grpcServer, &server{})
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println("nie działa")
	}
	log.Println("waiting for requests")
}
