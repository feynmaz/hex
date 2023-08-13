package main

import (
	"log"
	"os"

	"github.com/feynmaz/hex/internal/adapters/app/api"
	"github.com/feynmaz/hex/internal/adapters/core/arithmetic"

	"github.com/feynmaz/hex/internal/adapters/framework/left/grpc"
	"github.com/feynmaz/hex/internal/adapters/framework/right/db"
	"github.com/feynmaz/hex/internal/ports"
)

func main() {
	var err error

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
	gRPCAdapter = grpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
