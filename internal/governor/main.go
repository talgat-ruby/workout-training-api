package governor

import (
	"context"
	"log/slog"
	"workout-training-api/internal/config"
	"workout-training-api/internal/grpc/interceptor"
	"workout-training-api/internal/types/database"
)

type Governor struct {
	//Interceptor

	//Auth

	//Expense

}

func New(conf *config.Config) *Governor {
	return &Governor{
		//
		//
		//
	}
}

func (g *Governor) Config(ctx context.Context, conf *config.Config, logger *slog.Logger, db database.Database) {
	*g.Interceptor = *interceptor.New(conf, logger.With(slog.String("module", "interceptor")))
	*g.Auth = *auth.New(conf, logger.With(slog.String("component", "auth")), db)
	//*g.Expense = *expense.New(conf, logger.With(slog.String("component", "expense")), db)
}
