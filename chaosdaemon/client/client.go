package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "entropie.ai/carnot/chaosdaemon/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChaosDaemonClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	target := &pb.Target{Port: "31006"}
	stream, err := c.CaptureTraffic(ctx, target)
	for {
		payload, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Payload: %s", payload.GetBody())
	}
}
