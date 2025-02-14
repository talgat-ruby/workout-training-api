package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"workout-training-api/internal/config"
	pingSrvc "workout-training-api/internal/grpc/generated/workout-training-api/ping/v1"
	workoutv1 "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
	"workout-training-api/internal/grpc/interceptor"
	"workout-training-api/internal/grpc/ping"
	"workout-training-api/internal/grpc/workout"
	"workout-training-api/internal/types/controller"
)

type Grpc struct {
	conf   *config.APIGrpcConfig
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(config *config.APIGrpcConfig, logger *slog.Logger, ctrl controller.Controller) *Grpc {
	return &Grpc{
		conf:   config,
		logger: logger,
		ctrl:   ctrl,
	}
}

func (g *Grpc) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", g.conf.Host, g.conf.Port))
	if err != nil {
		g.logger.ErrorContext(ctx, "failed to listen", slog.Any("error", err), slog.Int("port", g.conf.Port))
		return err
	}

	inter := interceptor.New(g.logger, g.ctrl)
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(inter.Unary),
		grpc.StreamInterceptor(inter.Stream),
	)

	pingSrvc.RegisterPingServiceServer(
		srv,
		ping.New(
			g.logger.With(slog.String("component", "sanitary")),
		),
	)

	workoutv1.RegisterWorkoutServiceServer(
		srv,
		workout.New(
			g.logger.With(slog.String("component", "workout")),
			g.ctrl,
		),
	)

	// Register reflection service on gRPC server.
	reflection.Register(srv)

	g.logger.InfoContext(ctx, "starting server", slog.Int("port", g.conf.Port))
	if err := srv.Serve(lis); err != nil {
		g.logger.ErrorContext(ctx, "failed to serve", slog.Any("error", err), slog.Int("port", g.conf.Port))
		return err
	}

	return nil
}
