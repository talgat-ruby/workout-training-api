package workout

import (
	"context"
	"fmt"
	"workout-training-api/internal/types/controller"
)

func (w *WorkoutController) UpdateWorkout(ctx context.Context, req controller.UpdateWorkoutReq) (controller.UpdateWorkoutResp, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	exercises := make([]string, 0, len(req.GetExercises()))
	for _, e := range req.GetExercises() {
		exercises = append(exercises, e.GetID())
	}

	workout := &Workout{
		ID:        req.GetID(),
		UserID:    userID,
		Exercises: exercises,
		Comments:  req.GetComments(),
	}

	_, err := w.db.UpdateWorkout(ctx, workout)
	if err != nil {
		return nil, fmt.Errorf("failed to update workout: %w", err)
	}

	return &updateWorkoutResponse{}, nil
}

type updateWorkoutResponse struct{}
