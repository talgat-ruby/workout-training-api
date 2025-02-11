package database

import (
	"context"
	"time"
)

// Интерфейс запроса отчета по тренировкам
type GenerateWorkoutReportsReq interface {
	GetUserID() string
}

// Интерфейс ответа отчета по тренировкам
type GenerateWorkoutReportsRes interface {
	GetProgress() string
	GetWorkouts() []WorkoutResp
}

// Интерфейс для Workout
type WorkoutResp interface {
	GetID() string
	GetUserID() string
	GetExerciseIDs() []string
	GetScheduledTimes() []time.Time
}

// Внутренняя структура данных для сканирования из базы
type workoutData struct {
	ID             string
	UserID         string
	ExerciseIDs    []string
	ScheduledTimes []time.Time
}

// Реализация методов для WorkoutResp
func (w workoutData) GetID() string                  { return w.ID }
func (w workoutData) GetUserID() string              { return w.UserID }
func (w workoutData) GetExerciseIDs() []string       { return w.ExerciseIDs }
func (w workoutData) GetScheduledTimes() []time.Time { return w.ScheduledTimes }

// Реализация ответа для GenerateWorkoutReportsRes
type generateWorkoutReportsRes struct {
	Progress string
	Workouts []WorkoutResp
}

func (r generateWorkoutReportsRes) GetProgress() string {
	return r.Progress
}

func (r generateWorkoutReportsRes) GetWorkouts() []WorkoutResp {
	return r.Workouts
}

// Метод GenerateWorkoutReports теперь работает через Database
func (d *Database) GenerateWorkoutReports(ctx context.Context, req GenerateWorkoutReportsReq) (GenerateWorkoutReportsRes, error) {
	rows, err := d.DB.QueryContext(ctx, `
		SELECT id, user_id, exercise_ids, scheduled_times FROM workouts WHERE user_id = $1`, req.GetUserID())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workouts []WorkoutResp
	for rows.Next() {
		var w workoutData
		err := rows.Scan(&w.ID, &w.UserID, &w.ExerciseIDs, &w.ScheduledTimes)
		if err != nil {
			return nil, err
		}
		workouts = append(workouts, w)
	}

	// Возвращаем структуру, реализующую интерфейс
	return generateWorkoutReportsRes{
		Progress: "In Progress", // Можно добавить реальную логику расчета прогресса
		Workouts: workouts,
	}, nil
}
