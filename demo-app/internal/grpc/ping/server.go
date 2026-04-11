package ping

import (
	"context"
	"log"
	"time"

	"github.com/fallra1n/demo/proto/gen/go/ping"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ping.UnimplementedPingServer
}

func Register(gRPC *grpc.Server) {
	ping.RegisterPingServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Ping(ctx context.Context, req *ping.Request) (*ping.Response, error) {
	const op = "ping.serverAPI.Ping"

	log.Printf("%s: received request: %v", op, req)

	time.Sleep(100 * time.Millisecond)

	return &ping.Response{
		Message: "Pong: " + req.GetMessage(),
	}, nil
}
