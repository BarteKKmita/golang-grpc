package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not created connection")
	}
	defer connection.Close()

	client := calculatorpb.NewCalculatorServiceClient(connection)
	callServerStreaming(client)
}

func callUnary(client calculatorpb.CalculatorServiceClient) {
	sum, err := client.Sum(context.TODO(), &calculatorpb.SumRequest{FirstNumber: 2, SecondNumber: 4})
	if err != nil {
		log.Println("sth went wrong", err)
	}
	log.Println(sum.SumResult)
}

func callServerStreaming(client calculatorpb.CalculatorServiceClient) {

	decomposition, err := client.PrimeNumberDecomposition(context.TODO(), &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 94596618,
	})
	if err != nil {
		log.Printf("Server returned error %v\n", err)
	}
	for {
		rec, err := decomposition.Recv()
		if err == io.EOF {
			log.Println("retuning")
			break
		}
		if err != nil {
			return
		}
		log.Println(rec)
	}

}
