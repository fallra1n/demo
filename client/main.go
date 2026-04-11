package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fallra1n/demo/proto/gen/go/ping"
)

func main() {
	// Tracer
	shutdown, err := InitTracer("demo-app-client")
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(context.Background())

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := ping.NewPingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Ping(ctx, &ping.Request{
		Message: "ping",
	})
	if err != nil {
		log.Fatalf("Ошибка при вызове метода: %v", err)
	}

	fmt.Printf("Ответ от сервера: %s\n", resp.GetMessage())
}
