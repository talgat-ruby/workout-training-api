package database

import "context"

type CreateExerciseReq interface {
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
}

type CreateExerciseResp interface{}

func (d *Database) CreateExercise(ctx context.Context, req CreateExerciseReq) (CreateWorkoutResp, error) {
	_, err := d.DB.ExecContext(ctx, `
		INSERT INTO exercises (name, description, categories, muscle_groups)
		VALUES ($1, $2, $3, $4)`,
		req.GetName(), req.GetDescription(), req.GetCategories(), req.GetMuscleGroups(),
	)
	if err != nil {
		return nil, err
	}

	return struct{}{}, nil
}
