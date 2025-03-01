package database

import (

	"context"
	"time"
	"workout-training-api/internal/constant"
	"workout-training-api/internal/postgres/db_types/workout"

)

type CreateWorkoutReq interface {
	GetUserID() string
	GetName() string
	GetExercises() []workout.Exercise
	GetStatus() constant.WorkoutStatus
	GetDescription() string
	GetScheduledDate() []time.Time
}


type CreateWorkoutResp interface{}

func (d *Database) CreateWorkout(ctx context.Context, req CreateWorkoutReq) (CreateWorkoutResp, error) {
	_, err := d.DB.ExecContext(ctx, `
		INSERT INTO workouts (user_id, exercise_ids, scheduled_times)
		VALUES ($1, $2, $3)`,
		req.GetUserID(), req.GetExerciseIDs(), req.GetScheduledTimes(),
	)
	if err != nil {
		return nil, err
	}

	return struct{}{}, nil

type CreateWorkoutResp interface {
	GetWorkout() workout.Workout

}

/*

type Workout struct {
	WorkoutID     string                 `db:"workout_id" json:"workout_id"`
	UserID        string                 `db:"user_id" json:"user_id"`
	Name          string                 `db:"name" json:"name"`
	Description   string                 `db:"description" json:"description,omitempty"`
	Exercises     []Exercise             `json:"exercises,omitempty"` // Loaded separately
	Status        constant.WorkoutStatus `db:"status" json:"status"`
	Comments      []Comment              `json:"comments,omitempty"` // Loaded separately
	ScheduledDate []*time.Time           `db:"scheduled_date" json:"scheduled_date,omitempty"`
	CreatedAt     *time.Time             `db:"created_at" json:"created_at"`
	UpdatedAt     *time.Time             `db:"updated_at" json:"updated_at"`
}
*/
