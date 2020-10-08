package chat

import (
	"context"
	"errors"

	"github.com/meghashyamc/letschat-server/cache"
)

const Port = "9000"
const messages = "messages"

type Server struct {
}

func (s *Server) SendMessage(ctx context.Context, in *Message) (*Message, error) {
	if in == nil {
		return in, errors.New("nil message sent to server")
	}
	if _, err := cache.LPush(in.GetId(), in.GetBody()); err != nil {
		return in, err
	}
	in.Delivered = true
	return in, nil
}

func (s *Server) ReadMessage(ctx context.Context, in *Message) (*Message, error) {
	if in == nil {
		return in, errors.New("nil message sent to server")
	}
	val, err := cache.RPop(in.GetId())
	if err != nil {
		return in, err
	}
	in.Body = val
	return in, nil

}
