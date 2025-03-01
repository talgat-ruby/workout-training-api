package request

import (
	"time"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/graphql/graph/model"
	"workout-training-api/internal/types/controller"
)

type CtrlCreateWorkoutRequest struct {
	name          string
	description   string
	exercises     []*model.ExerciseInput
	status        string
	scheduledDate []time.Time
}

func (c CtrlCreateWorkoutRequest) ScheduledDate() time.Time {
	return c.scheduledDate[0]
}

func (c CtrlCreateWorkoutRequest) GetName() string {
	return c.name
}

func (c CtrlCreateWorkoutRequest) GetDescription() string {
	return c.description
}

func (c CtrlCreateWorkoutRequest) GetExercises() []controller.Exercise {
	if c.exercises == nil {
		return nil
	}

	exercises := make([]controller.Exercise, len(c.exercises))
	for i, ex := range c.exercises {
		if ex == nil { // Handle nil pointers
			continue
		}
		exercises[i] = ExerciseImpl{ // Use the new struct
			ID:          ex.ID,
			Name:        ex.Name,
			Description: ex.Description,
			MuscleGroup: ex.MuscleGroup,
			Category:    ex.Category,
			Repetitions: int(ex.Repetitions),
			Sets:        int(ex.Sets),
			Weight:      float64(ex.Weight),
		}
	}
	return exercises
}

type ExerciseImpl struct {
	ID          string
	Name        string
	Description string
	MuscleGroup string
	Category    string
	Repetitions int
	Sets        int
	Weight      float64
}

func (e ExerciseImpl) GetID() string          { return e.ID }
func (e ExerciseImpl) GetName() string        { return e.Name }
func (e ExerciseImpl) GetDescription() string { return e.Description }
func (e ExerciseImpl) GetMuscleGroup() string { return e.MuscleGroup }
func (e ExerciseImpl) GetCategory() string    { return e.Category }
func (e ExerciseImpl) GetRepetitions() int    { return e.Repetitions }
func (e ExerciseImpl) GetSets() int           { return e.Sets }
func (e ExerciseImpl) GetWeight() float64     { return e.Weight }

func (c CtrlCreateWorkoutRequest) GetStatus() constant.WorkoutStatus {
	return constant.WorkoutStatus(c.status)
}

type CtrlUpdateWorkoutRequest struct {
	ID        string
	Exercises []*model.ExerciseInput
	Comments  []string
}

func (c CtrlUpdateWorkoutRequest) GetExercises() []controller.Exercise {
	//TODO implement me
	panic("implement me")
}

func (c CtrlUpdateWorkoutRequest) GetID() string {
	return c.ID
}

func (c CtrlUpdateWorkoutRequest) GetComments() []string {
	return c.Comments
}

type CtrlListWorkoutRequest struct {
	UserID string
}

func (c CtrlListWorkoutRequest) GetUserID() string {
	return c.UserID
}

type CtrlDeleteWorkoutRequest struct {
	ID string
}

func (c CtrlDeleteWorkoutRequest) GetID() string {
	return c.ID
}

func NewCtrlCreateWorkoutRequest(name string, description string, exercises []*model.ExerciseInput, status string, scheduledDate []time.Time) *CtrlCreateWorkoutRequest {
	return &CtrlCreateWorkoutRequest{
		name: name, description: description, exercises: exercises, status: status, scheduledDate: scheduledDate,
	}
}

func NewCtrlDeleteWorkoutRequest(workoutID string) *CtrlDeleteWorkoutRequest {
	return &CtrlDeleteWorkoutRequest{
		ID: workoutID,
	}
}

func NewCtrlUpdateWorkoutRequest(workout *model.WorkoutInput) *CtrlUpdateWorkoutRequest {
	return &CtrlUpdateWorkoutRequest{
		ID:        workout.ID,
		Exercises: workout.Exercises,
		Comments:  nil,
	}
}

func NewCtrlListWorkoutsRequest(userID string) *CtrlListWorkoutRequest {
	return &CtrlListWorkoutRequest{
		UserID: userID,
	}
}
