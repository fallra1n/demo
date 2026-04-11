package grpc

import (
	"context"
	"log"
	"net"
	"time"

	pinggrpc "github.com/fallra1n/demo/demo-app/internal/grpc/ping"
	otelLib "github.com/fallra1n/demo/demo-app/internal/lib/otel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer  *grpc.Server
	tracerClose func(context.Context) error
}

func NewApp() *App {
	// Tracer
	tracerClose, err := otelLib.InitTracer("demo-app")
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}

	a := &App{
		gRPCServer: grpc.NewServer(
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
		),
		tracerClose: tracerClose,
	}

	a.registerServices()
	return a
}

func (a *App) Run() error {
	log.Println("Starting gRPC server...")

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (a *App) registerServices() {
	pinggrpc.Register(a.gRPCServer)
}

func (a *App) Close() error {
	ctx, err := context.WithTimeout(context.Background(), 5*time.Second)
	if err != nil {
		log.Printf("failed to create context for tracer shutdown: %v", err)
	}

	a.gRPCServer.GracefulStop()
	a.tracerClose(ctx)
	return nil
}
