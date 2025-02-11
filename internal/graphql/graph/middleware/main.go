package middleware

import (
	"log/slog"

	"workout-training-api/internal/types/controller"
)

type Middleware struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(
	logger *slog.Logger,
	ctrl controller.Controller,
) *Middleware {
	return &Middleware{
		logger: logger,
		ctrl:   ctrl,
	}
}
