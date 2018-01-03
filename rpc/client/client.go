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
	server := fmt.Sprintf("%v:1080", os.Args[1])
	fmt.Printf("Connecting to %v\n", server)
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetuserClient(conn)

	name := os.Args[2]
	for i := 0; i < 100; i++ {
		resp, err := client.GetGreeting(context.Background(), &pb.Person{
			Name: name,
			Age:  36,
		})
		if err != nil {
			log.Fatalf("Received an error from gRPC server: %v", err)
		}
		fmt.Printf("%s", resp.GetGreeting())
	}
}
