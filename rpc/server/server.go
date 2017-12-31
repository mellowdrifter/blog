package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mellowdrifter/blog/proto/greeting"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Listening on port 1080")
	lis, err := net.Listen("tcp", "localhost:1080")
	if err != nil {
		log.Fatalf("failed to bind to port: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterGreetuserServer(grpcServer, &server{})

	grpcServer.Serve(lis)

}

func (s *server) GetGreeting(ctx context.Context, req *pb.Person) (*pb.Greeting, error) {
	name := req.GetName()

	fmt.Printf("Received the name: %v\n", name)

	return &pb.Greeting{
		Greeting: fmt.Sprintf("Hello %v\n", name),
	}, nil

}
