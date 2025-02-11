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

type ExerciseModel struct {
	ID           uint `gorm:"primaryKey"`
	WorkoutID    uint `gorm:"index"`
	Name         string
	Description  string
	Categories   []string `gorm:"type:text[]"`
	MuscleGroups []string `gorm:"type:text[]"`
	Reps         int
	Sets         int
}

func CreateExercise(ctx context.Context, wourkoutID uint, req CreateExerciseReq) (*ExerciseModel, error) {
	exercise := ExerciseModel{
		WorkoutID:    wourkoutID,
		Name:         req.GetName(),
		Description:  req.GetDescription(),
		Categories:   req.GetCategories(),
		MuscleGroups: req.GetMuscleGroups(),
	}
	if err := DB.WithContext(ctx).Create(&exercise).Error; err != nil {
		return nil, err
	}
	return &exercise, nil
}
