package workout

import (
	"context"

	"workout-training-api/internal/types/controller"
)

type listWorkoutsResponse struct {
	Workouts []Workout 
}

func (w *WorkoutController) ListWorkouts(ctx context.Context, req controller.ListWorkoutsReq) (controller.ListWorkoutsResp, error) {
	panic("implement me")
}

func (r *listWorkoutsResponse) GetWorkouts() []Workout {
	return r.Workouts
}