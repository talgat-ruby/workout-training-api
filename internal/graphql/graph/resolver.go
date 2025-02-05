package graph

import (
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/types/controller"
)

type Resolver struct {
	conf   *config.APIGraphQLConfig
	ctrl   controller.Controller
	logger *slog.Logger
}

func NewResolver(conf *config.APIGraphQLConfig, ctrl controller.Controller, logger *slog.Logger) *Resolver {
	return &Resolver{
		conf:   conf,
		ctrl:   ctrl,
		logger: logger.With(slog.String("module", "resolver")),
	}
}
