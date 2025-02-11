package directives

import (
	"log/slog"

	"workout-training-api/internal/types/controller"
)

type Directives struct {
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(logger *slog.Logger, ctrl controller.Controller) *Directives {
	return &Directives{
		logger: logger.With(slog.String("module", "directives")),
		ctrl:   ctrl,
	}
}
