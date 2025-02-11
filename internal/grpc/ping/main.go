package ping

import (
	"log/slog"
	pingSrvc "workout-training-api/internal/grpc/generated/workout-training-api/ping/v1"
)

type Ping struct {
	pingSrvc.UnimplementedPingServiceServer
	log *slog.Logger
}

func New(log *slog.Logger) *Ping {
	return &Ping{
		log: log,
	}
}
