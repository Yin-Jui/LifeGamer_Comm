package main

import(
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../proto"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server (implement comm engine server)
type server struct{}

func (s *server) RawData(ctx context.Context, in *pb.RawRequest) (*pb.RawReply, error) {
	return &pb.RawReply{Data: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}