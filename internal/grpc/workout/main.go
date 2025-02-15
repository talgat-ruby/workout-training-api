package workout

import (
	"log/slog"
	workoutSrvc "workout-training-api/internal/grpc/generated/workout-training-api/workout/v1"
	controller "workout-training-api/internal/types/controller"
)

type Workout struct {
	workoutSrvc.UnimplementedWorkoutServiceServer
	log  *slog.Logger
	ctrl controller.Controller
}

func New(log *slog.Logger, ctrl controller.Controller) *Workout {
	return &Workout{
		log:  log,
		ctrl: ctrl,
	}
}
