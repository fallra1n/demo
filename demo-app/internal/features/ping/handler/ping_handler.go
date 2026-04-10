package handler

type pingHandler struct {
}

func NewPingHandler() *pingHandler {
	return &pingHandler{}
}

func (h *pingHandler) Handle() string {
	return "pong"
}
