package graph

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"workout-training-api/internal/types/controller"
)

type GraphQL struct {
	conf   *config.APIGraphQLConfig
	logger *slog.Logger
	ctrl   contr 1 	oller.Controller
}

func New(conf *config.APIGraphQLConfig, logger *slog.Logger, ctrl controller.Controller) *GraphQL {
	return &GraphQL{
		conf:   conf,
		logger: logger,
		ctrl:   ctrl,
	}
}

func (a *GraphQL) Start(ctx context.Context) error {
	midd := middleware.New(a.logger, a.ctrl)
	d := directives.New(a.logger, a.ctrl)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(a.conf, a.ctrl, a.logger),
				Directives: graph.DirectiveRoot{
					Auth: d.AuthDirective,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", midd.Authenticator(srv))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", a.conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.conf.Port), nil))

	return nil
}
