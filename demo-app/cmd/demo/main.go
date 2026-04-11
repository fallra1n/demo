package main

import (
	"os"
	"syscall"

	appgrpc "github.com/fallra1n/demo/demo-app/internal/app/grpc"
	"github.com/fallra1n/demo/demo-app/internal/lib/shutdown"
)

func main() {
	app := appgrpc.NewApp()
	go app.Run()

	shutdown.Graceful([]os.Signal{syscall.SIGINT, syscall.SIGTERM}, app)
}
