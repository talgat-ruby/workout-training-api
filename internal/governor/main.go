package governor

import (
	"context"
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/governor/auth"
	"workout-training-api/internal/governor/interceptor"
	"workout-training-api/internal/governor/workout"
	"workout-training-api/internal/types/database"
)

type Governor struct {
	*interceptor.Interceptor
	*auth.Auth
	*workout.Workout
}

func New(conf *config.Config) *Governor {
	return &Governor{
		Interceptor: new(interceptor.Interceptor),
		Auth:        new(auth.Auth),
		Workout:     new(workout.Workout),
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	*g.Interceptor = *interceptor.New(conf, logger.With(slog.String("module", "interceptor")))
	*g.Auth = *auth.New(conf, logger.With(slog.String("component", "auth")), db)
	*g.Workout = *workout.New(conf, logger.With(slog.String("component", "workout")), db)
}
