package database

import "context"

type DeleteWorkoutReq interface {
	GetUserID() string
	GetID() string
}

type DeleteWorkoutResp interface{}

func (d *Database) DeleteWorkout(ctx context.Context, req DeleteWorkoutReq) (DeleteWorkoutResp, error) {
	_, err := d.DB.ExecContext(ctx, `DELETE FROM workouts WHERE id = $1 AND user_id = $2`, req.GetID(), req.GetUserID())
	if err != nil {
		return nil, err
	}

	return struct{}{}, nil
}
