package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
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

	doUnary(err, client)
}

func doUnary(err error, client greetpb.GreetServiceClient) {
	greetRequest := greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: "Bartek", LastName: "DUUUPA"}}
	greet, err := client.Greet(context.TODO(), &greetRequest)
	if err != nil {
		log.Print("Coś poszło nie tak")

		log.Println(err)
	}
	log.Println(greet.Result)
}
