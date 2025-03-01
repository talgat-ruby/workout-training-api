package workout

import (
	"context"
	"fmt"
	"time"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"
	"workout-training-api/internal/types/controller"
)

type WorkoutStatus string

func (w *WorkoutController) CreateWorkout(ctx context.Context, req controller.CreateWorkoutReq) (controller.CreateWorkoutResp, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		return nil, fmt.Errorf("user not authenticated")
	}
	exercises := make([]workout.Exercise, 0, len(req.GetExercises()))

	for _, e := range req.GetExercises() {
		exercises = append(exercises, workout.Exercise{
			ExerciseID:   e.GetID(),   
			Name: e.GetName(), 
		})
	}
	workout := &Workout{
		UserID:        userID,
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Status:        req.GetStatus(),
		Exercises:     exercises,
		ScheduledDate: req.ScheduledDate(),
	}

	_, err := w.db.CreateWorkout(ctx, workout)
	if err != nil {
		return nil, fmt.Errorf("failed to save workout: %w", err)
	}

	return &createWorkoutResponse{}, nil
}

type createWorkoutResponse struct {
	ID     string
	Status constant.WorkoutStatus
}

type Workout struct {
	ID            string
	UserID        string
	Name          string
	Description   string
	Status        constant.WorkoutStatus
	Exercises     []workout.Exercise
	Comments      []string
	ScheduledDate time.Time
}

func (w *Workout) GetUserID() string {
	return w.UserID
}

func (w *Workout) GetExerciseIDs() []workout.Exercise {
	exerciseIDs := make([]string, 0, len(w.Exercises))
	for _, e := range w.Exercises {
		exerciseIDs = append(exerciseIDs, e.ExerciseID)
	}
	return w.Exercises
}

func (w *Workout) GetID() string {
	return w.ID
}

func (w *Workout) GetComments() []string {
	return w.Comments
}

func (w *Workout) GetDescription() string {
	return w.Description
}

func (w *Workout) GetExercises() []workout.Exercise {
	return w.Exercises
}

func (w *Workout) GetName() string {
	return w.Name
}

func (w *Workout) GetScheduledDate() []time.Time {
	return []time.Time{w.ScheduledDate}
}

func (w *Workout) GetStatus() constant.WorkoutStatus {
	return w.Status
}

func (r *createWorkoutResponse) GetID() string {
	return r.ID
}

func (r *createWorkoutResponse) GetStatus() constant.WorkoutStatus {
	return r.Status
}
