package request

import (
	"time"
	"workout-training-api/internal/graphql/graph/model"
)

type CtrlCreateWorkoutRequest struct {
	name          string
	description   string
	exercises     []*model.ExerciseInput
	scheduledTime time.Time
}

type CtrlUpdateWorkoutRequest struct {
	ID            string
	name          string
	description   string
	exercises     []*model.ExerciseInput
	scheduledTime int32
}

func NewCtrlCreateWorkoutRequest(name string, description string, exercises []*model.ExerciseInput, scheduledTime time.Time) *CtrlCreateWorkoutRequest {
	return &CtrlCreateWorkoutRequest{
		name: name, description: description, exercises: exercises, scheduledTime: scheduledTime,
	}
}

func NewCtrlDeleteWorkoutRequest(workoutID string) string {
	return workoutID
}

func NewCtrlUpdateWorkoutRequest(workout *model.WorkoutInput) *CtrlUpdateWorkoutRequest {
	return &CtrlUpdateWorkoutRequest{
		ID:            workout.ID,
		name:          workout.Name,
		description:   workout.Description,
		exercises:     workout.Exercises,
		scheduledTime: workout.ScheduledTimes,
	}
}
