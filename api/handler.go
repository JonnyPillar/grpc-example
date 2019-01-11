package api

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Println("Messaged Reseived", in.Greeting)
	return &PingMessage{Greeting: "Bar"}, nil
}
