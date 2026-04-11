package grpc

import (
	"log"
	"net"

	pinggrpc "github.com/fallra1n/demo/demo-app/internal/grpc/ping"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
}

func NewApp() *App {
	a := &App{
		gRPCServer: grpc.NewServer(),
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
	a.gRPCServer.GracefulStop()
	return nil
}
