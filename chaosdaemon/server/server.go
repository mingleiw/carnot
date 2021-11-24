package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "entropie.ai/carnot/chaosdaemon/pb"
	"entropie.ai/carnot/pkg/capture"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type chaosDaemonServer struct {
	pb.UnimplementedChaosDaemonServer
}

func (s *chaosDaemonServer) CaptureTraffic(target *pb.Target, stream pb.ChaosDaemon_CaptureTrafficServer) error {

	c := &capture.Capture{}
	go c.WithIface("lo").WithPort(target.Port).Start()
	payload := &pb.Payload{Body: "test body"}
	stream.Send(payload)
	return nil
}

func newServer() *chaosDaemonServer {
	s := &chaosDaemonServer{}

	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterChaosDaemonServer(grpcServer, newServer())

	log.Println("starting server...")
	grpcServer.Serve(lis)
}
