package main

import (
	"context"
	"log/slog"
	askar_postgres "workout-training-api/internal/askar-postgres"
	"workout-training-api/internal/config"
	"workout-training-api/internal/constant"
	"workout-training-api/pkg/logger"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())

	// // config
	// conf := config.New(ctx)

	// // logger
	// log := logger.New(conf.ENV != constant.EnvironmentLocal)

	// p, err := postgres.New(conf.Postgres, log.With(slog.String("service", "postgre")))
	// if err != nil {
	// 	log.ErrorContext(ctx, "failed to start postgre", slog.Any("error", err))
	// 	panic(err)
	// }

	// ctrl := controller.New(conf, p)

	// gr := grpc.New(conf.API.Grpc, log.With(slog.String("service", "grpc")), ctrl)
	// go func(ctx context.Context, cancelFunc context.CancelFunc) {
	// 	if err := gr.Start(ctx); err != nil {
	// 		log.ErrorContext(ctx, "failed to start grpc", slog.Any("error", err))
	// 	}

	// 	cancelFunc()
	// }(ctx, cancel)

	// gql := graphql.New(conf.API.GraphQL, log.With(slog.String("service", "graphql")), ctrl)
	// go func(ctx context.Context, cancelFunc context.CancelFunc) {
	// 	if err := gql.Start(ctx); err != nil {
	// 		log.ErrorContext(ctx, "failed to start graphql", slog.Any("error", err))
	// 	}

	// 	cancelFunc()
	// }(ctx, cancel)

	// go func(cancelFunc context.CancelFunc) {
	// 	shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
	// 	signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

	// 	sig := <-shutdown
	// 	log.WarnContext(ctx, "signal received - shutting down...", slog.Any("signal", sig))

	// 	cancelFunc()
	// }(cancel)

	// <-ctx.Done()

	// fmt.Println("shutting down gracefully")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(conf.ENV != constant.EnvironmentLocal)

	// postgres
	_, err := askar_postgres.New(conf.AskarPostgres, log)
	if err != nil {
		log.ErrorContext(ctx, "failed to start postgres", slog.Any("error", err))
	}
	//fmt.Print(db)
	log.InfoContext(ctx, "postgres connection pool established", nil)

}
