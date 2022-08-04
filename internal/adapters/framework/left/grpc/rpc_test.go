package grpc

import (
	"log"
	"os"

	"context"
	"github.com/PrinceDavis/hex/internal/adapters/app/api"
	"github.com/PrinceDavis/hex/internal/adapters/core/arithmetic"
	"github.com/PrinceDavis/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/PrinceDavis/hex/internal/adapters/framework/right/db"
	"github.com/PrinceDavis/hex/internal/ports"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error

	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

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
	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start test server: %v", err)
		}
	}()
}

func bufDailer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDailer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}
	return conn
}

func TestAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := pb.OperationParameters{
		A: 1,
		B: 2,
	}

	answer, err := client.GetAddition(ctx, &params)
	if err != nil {
		t.Fatalf("expected: %v,  got: %v", nil, err)
	}
	require.Equal(t, int32(3), answer.Value)
}
