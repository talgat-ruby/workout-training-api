package ping

import (
	"log/slog"
)

type Ping struct {
	pingv1.Unimmplemen
	log *slog.Logger
}

func New(log *slog.Logger) *Sanitary {
	return &Sanitary{
		log: log,
	}
}
