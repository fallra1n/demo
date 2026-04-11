package ping

import (
	"context"

	"github.com/fallra1n/demo/proto/gen/go/ping"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ping.UnimplementedPingServer
}

func Register(gRPC *grpc.Server) {
	ping.RegisterPingServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Do(ctx context.Context, req *ping.Request) (*ping.Response, error) {
	return &ping.Response{
		Message: "Pong: " + req.GetMessage(),
	}, nil
}
