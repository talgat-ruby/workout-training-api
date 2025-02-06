package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"workout-training-api/internal/config"
	"workout-training-api/internal/grpc/interceptor"
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

	sanitaryv1.RegisterSanitaryServiceServer(
		srv,
		sanitary.New(
			g.logger.With(slog.String("component", "sanitary")),
		),
	)
	expensev1.RegisterExpenseServiceServer(
		srv,
		expense.New(
			g.logger.With(slog.String("component", "expense")),
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
