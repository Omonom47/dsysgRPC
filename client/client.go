package main

import (
	"context"
	"fmt"
	"log"
	"os"

	t "time"

	p "github.com/Omonom47/dsysgRPC/proto"

	"google.golang.org/grpc"
)

func main() {

	go TimeServiceClient()
	go TimeServiceClient()

	for {

	}
}

func TimeServiceClient(ip string) {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(ip+":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := p.NewGetCurrentTimeClient(conn)

	for {
		SendGetTimeRequest(c)
		t.Sleep(5 * t.Second)
	}
}

func SendGetTimeRequest(c p.GetCurrentTimeClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := p.GetTimeRequest{}

	response, err := c.GetTime(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	fmt.Printf("Current time right now: %s\n", response.Reply)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println(response.Reply)
}
