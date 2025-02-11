package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"workout-training-api/internal/config"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/controller"
	"workout-training-api/internal/governor"
	"workout-training-api/internal/grpc"
	"workout-training-api/internal/postgres"
	"workout-training-api/pkg/logger"
)

<<<<<<< Updated upstream

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"workout-training-api/internal/config"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/governor"
	"workout-training-api/internal/grpc"
	"workout-training-api/internal/postgres"
	"workout-training-api/pkg/logger"
)


func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// config
	conf := config.New(ctx)

	// logger
	log := logger.New(conf.ENV != constant.EnvironmentLocal)

	p, err := postgres.New(conf.Postgres, log.With(slog.String("service", "postgre")))
	if err != nil {
		log.ErrorContext(ctx, "failed to start postgre", slog.Any("error", err))
		panic(err)
	}

	ctrl := controller.New(conf, p)

	gr := grpc.New(conf.API.Grpc, log.With(slog.String("service", "grpc")), ctrl)
	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		if err := gr.Start(ctx); err != nil {
			log.ErrorContext(ctx, "failed to start grpc", slog.Any("error", err))
		}

		cancelFunc()
	}(ctx, cancel)

	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		log.WarnContext(ctx, "signal received - shutting down...", slog.Any("signal", sig))

		cancelFunc()
	}(cancel)

	<-ctx.Done()

	fmt.Println("shutting down gracefully")
}
