package main

import (
	"log"
	"os"

	"github.com/PrinceDavis/hex/internal/adapters/app/api"
	"github.com/PrinceDavis/hex/internal/adapters/core/arithmetic"
	gRPC "github.com/PrinceDavis/hex/internal/adapters/framework/left/grpc"
	"github.com/PrinceDavis/hex/internal/adapters/framework/right/db"
	"github.com/PrinceDavis/hex/internal/ports"
)

func main() {
	var err error

	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsource := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsource)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection %v", err)
	}
	defer dbaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbaseAdapter)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
