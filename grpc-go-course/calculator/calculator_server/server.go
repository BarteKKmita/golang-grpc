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
		log.Println("failed to create serve listener")
	}
	log.Println("waiting for requests")
}

func (s *server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("About to handle sum request")
	sumResponse := calculatorpb.SumResponse{
		SumResult: request.FirstNumber + request.SecondNumber,
	}
	return &sumResponse, nil
}

func (s *server) PrimeNumberDecomposition(request *calculatorpb.PrimeNumberDecompositionRequest,
	svr calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Received PrimeNumberDecomposition request %v\n", request)
	number := request.GetNumber()
	var divisor int64 = 2
	for number > 1 {
		if number%divisor == 0 {
			err := svr.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			if err != nil {
				log.Println("Failed to send resposne")
			}
			number = number / divisor
		} else {
			divisor += 1
		}
	}
	return nil
}
