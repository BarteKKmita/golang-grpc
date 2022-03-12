package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"log"
)

func main() {
	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not created connection")
	}
	defer connection.Close()

	client := calculatorpb.NewCalculatorServiceClient(connection)
	sum, err := client.Sum(context.TODO(), &calculatorpb.SumRequest{FirstNumber: 2, SecondNumber: 4})
	if err != nil {
		log.Println("sth went wrong", err)
	}
	log.Println(sum.SumResult)
}
