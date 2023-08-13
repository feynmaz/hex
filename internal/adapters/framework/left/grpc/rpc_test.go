package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/feynmaz/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/stretchr/testify/require"

	"github.com/feynmaz/hex/internal/adapters/app/api"
	"github.com/feynmaz/hex/internal/adapters/core/arithmetic"
	"github.com/feynmaz/hex/internal/adapters/framework/right/db"
	"github.com/feynmaz/hex/internal/ports"
	"google.golang.org/grpc"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis := bufconn.Listen(bufSize)
	grpcServer := gogrpc.NewServer()

	// ports
	var dbAdapter ports.DB
	var core ports.Arithmetic
	var appAdapter ports.API
	var gRPCAdapter ports.GRPC

	// right framework
	dbDriver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("DSN")
	dbAdapter, err = db.NewAdapter(dbDriver, dsn)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbAdapter.CloseDBConnection()

	// domain
	core = arithmetic.NewAdapter()

	// application
	appAdapter = api.NewAdapter(dbAdapter, core)

	// left framework
	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start test server: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *gogrpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", gogrpc.WithContextDialer(bufDialer), gogrpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(2))
}
