package database

import (
	"context"
	"time"
)

type CreateWorkoutReq interface {
	GetUserID() string
	GetExerciseIDs() []string
	GetScheduledTimes() []time.Time
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
}
