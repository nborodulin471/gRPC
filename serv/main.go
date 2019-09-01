package main

import (
	"context"
	pb "github.com/nborodulin471/gRPC/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
}

func (*server) Send(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Recd: true}, nil
}

func (*server) Retrieve(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Started...")
	rpcserv := grpc.NewServer()
	pb.RegisterRequestReplyServer(rpcserv, &server{})
	reflection.Register(rpcserv)
	err = rpcserv.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Server is close")
}
