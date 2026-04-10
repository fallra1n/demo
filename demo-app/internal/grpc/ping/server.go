package ping

import "github.com/fallra1n/demo/proto/gen/go/ping"

type serverApi struct {
	ping.UnimplementedReverseServer
}