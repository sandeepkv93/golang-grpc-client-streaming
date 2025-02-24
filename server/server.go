package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-client-streaming/proto"
)

type server struct {
	pb.UnimplementedAverageServiceServer
}

func (s *server) ComputeAverage(stream pb.AverageService_ComputeAverageServer) error {
	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float32(sum) / float32(count)
			return stream.SendAndClose(&pb.AverageResponse{Average: average})
		}
		if err != nil {
			log.Printf("received error: %v", err)
			return err
		}

		num := req.GetNumber()
		log.Printf("Server received: %d", num)
		sum += int(num)
		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAverageServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
