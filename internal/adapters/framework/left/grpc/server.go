package grpc

import (
	"log"
	"net"

	"github.com/feynmaz/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/feynmaz/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.API
}

func NewAdapter(api ports.API) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()

	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
