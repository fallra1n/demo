package ping

import (
	"context"
	"log"
	"time"

	"github.com/fallra1n/demo/proto/gen/go/ping"
	"go.opentelemetry.io/otel/trace"
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

	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID().String()

	log.Printf("%s: received request: %v, traceID: %s", op, req, traceID)

	time.Sleep(100 * time.Millisecond)

	return &ping.Response{
		Message: "Pong: " + req.GetMessage(),
	}, nil
}
