package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fallra1n/demo/proto/gen/go/ping"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
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
