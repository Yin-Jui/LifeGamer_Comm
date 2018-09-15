package main

import(
	"log"
	"net"

	"fmt"
	"flag"

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

// ============================== RawData channel (Listening) ==============================
func (s *server) RawData(ctx context.Context, in *pb.RawRequest) (*pb.RawReply, error) {
	// reply msg if receive
	return &pb.RawReply{Data: "Hello " + in.Name}, nil
}
// ============================== RawData channel ==============================

func main() {
	// Add argparser for comm_server
	portN := flag.Int("port", 50051, "Port Number which listening by Comm Engine.")
	flag.Parse()

	// Create Server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCommServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	// NewServer listening on port 
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}