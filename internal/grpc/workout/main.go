package workout

import (
	"log/slog"
	pingSrvc "workout-training-api/internal/grpc/generated/workout-training-api/ping/v1"
	"workout-training-api/internal/types/controller"
)

// CreateWorkout implements the CreateWorkout gRPC method
type Workout struct {
	pingSrvc.UnimplementedPingServiceServer
	log  *slog.Logger
	ctrl controller.Controller
}

func New(log *slog.Logger, ctrl controller.Controller) *Workout {
	return &Workout{
		log:  log,
		ctrl: ctrl,
	}
}
