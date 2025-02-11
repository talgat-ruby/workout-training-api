package database

import (
	"context"
	"time"
)

// Интерфейс запроса списка упражнений
type ListExerciseReq interface{}

// Интерфейс ответа, содержащего список упражнений
type ListExerciseResp interface {
	GetList() []ExerciseResp
}

// Интерфейс одного упражнения
type ExerciseResp interface {
	GetID() string
	GetName() string
	GetDescription() string
	GetCategories() []string
	GetMuscleGroups() []string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

// Реализация `ExerciseResp`
type exerciseData struct {
	ID           string
	Name         string
	Description  string
	Categories   []string
	MuscleGroups []string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Методы для реализации `ExerciseResp`
func (e exerciseData) GetID() string             { return e.ID }
func (e exerciseData) GetName() string           { return e.Name }
func (e exerciseData) GetDescription() string    { return e.Description }
func (e exerciseData) GetCategories() []string   { return e.Categories }
func (e exerciseData) GetMuscleGroups() []string { return e.MuscleGroups }
func (e exerciseData) GetCreatedAt() time.Time   { return e.CreatedAt }
func (e exerciseData) GetUpdatedAt() time.Time   { return e.UpdatedAt }

// Ответ `ListExerciseResp`, содержащий список упражнений
type listExerciseResp struct {
	List []ExerciseResp
}

// Реализация метода `GetList()`
func (r listExerciseResp) GetList() []ExerciseResp {
	return r.List
}

// Метод получения списка упражнений
func (d *Database) ListExpense(ctx context.Context, req ListExerciseReq) (ListExerciseResp, error) {
	rows, err := d.DB.QueryContext(ctx, `
		SELECT id, name, description, categories, muscle_groups, created_at, updated_at FROM exercises`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []ExerciseResp
	for rows.Next() {
		var e exerciseData
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Categories, &e.MuscleGroups, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, e)
	}

	return listExerciseResp{List: exercises}, nil
}
