package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/mellowdrifter/blog/rpc/proto/greeting"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.20.30.30:1080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetuserClient(conn)

	name := os.Args[1]
	resp, err := client.GetGreeting(context.Background(), &pb.Person{Name: name})
	if err != nil {
		log.Fatalf("Received an error from gRPC server: %v", err)
	}
	fmt.Printf("%s", resp.GetGreeting())
}
