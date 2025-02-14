package workout

import (
	"context"
	"fmt"

	"workout-training-api/internal/types/controller"
)

func (w *WorkoutController) DeleteWorkout(ctx context.Context, req controller.DeleteWorkoutReq) (controller.DeleteWorkoutResp, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}

	workout := &Workout{
		ID:     req.GetID(),
		UserID: userID,
	}

	_, err := w.db.DeleteWorkout(ctx, workout)
	if err != nil {
		return nil, fmt.Errorf("failed to delete workout: %w", err)
	}

	return &deleteWorkoutResponse{}, nil
}

type deleteWorkoutResponse struct{}
