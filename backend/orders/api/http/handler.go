package http

import (
	"context"

	"github.com/mihailtudos/yumgo/common"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func Register(ctx context.Context, e common.EchoRouter, handler Handler) error {
	return nil
}
