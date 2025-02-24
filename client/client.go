package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-client-streaming/proto"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAverageServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.ComputeAverage(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Client streaming started: Send 10 numbers to the server for average calculation")
	for i := 1; i <= 10; i++ {
		req := &pb.AverageRequest{
			Number: int32(rand.Intn(1000) + 1),
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}
	fmt.Printf("Average: %f\n", reply.GetAverage())
}
