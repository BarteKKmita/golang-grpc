package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello I'am a client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()
	client := greetpb.NewGreetServiceClient(connection)
	ctx := context.TODO()
	getServerStream(client, ctx)
	//doUnary(err, client, &ctx)
}

func getServerStream(client greetpb.GreetServiceClient, ctx context.Context) {
	response, err := client.GreetManyTimes(ctx, &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Bartek",
			LastName:  "Kmita",
		},
	})
	if err != nil {
		log.Println("failed to create client")
	}
	for {
		recv, err := response.Recv()
		if err == io.EOF {
			log.Println("reached end of file")
			break
		}
		log.Println(recv.GetResult())
	}
}

func doUnary(client greetpb.GreetServiceClient, ctx *context.Context) {
	greetRequest := greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: "Bartek", LastName: "DUUUPA"}}
	greet, err := client.Greet(*ctx, &greetRequest)
	if err != nil {
		log.Print("server returned error")
		log.Println(err)
	}
	log.Println(greet.Result)
}
