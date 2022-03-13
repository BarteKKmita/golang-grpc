package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func main() {
	fmt.Println("Hello world")
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s server) Greet(ctx context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet was invoked with %v", request)
	firstName := request.Greeting.FirstName
	return &greetpb.GreetResponse{Result: "Hello" + firstName}, nil
}

func (s server) GreetManyTimes(request *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greet many times was invoked with %v", request)
	firstName := request.Greeting.FirstName
	for i := 0; i < 10; i++ {
		response := greetpb.GreetManyTimesResponse{
			Result: "Hello" + firstName + "for the" + strconv.Itoa(i) + "time",
		}
		err := stream.Send(&response)
		if err != nil {
			log.Printf("failed to send message %v\n", response.Result)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
