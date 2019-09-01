package main

import (
	"context"
	pb "github.com/nborodulin471/gRPC/transport"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewRequestReplyClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rply, err := c.Send(ctx, &pb.Request{To: "first", Text: "test"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rply.Recd)

	rply, err = c.Retrieve(ctx, &pb.Request{To: "second", Text: "test"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(rply.Recd)
}
