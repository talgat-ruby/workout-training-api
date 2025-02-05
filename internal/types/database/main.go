package database

import (
	"context"
	"database/sql"
	"fmt"
)

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUser(context.Context, FindUserReq) (FindUserResp, error)
}

type Exercise interface {
	CreateExercise(context.Context, CreateExerciseReq) (CreateWorkoutResp, error)
	UpdateExercise(context.Context, UpdateExerciseReq) (UpdateExerciseResp, error)
	DeleteExercise(context.Context, DeleteExerciseReq) (DeleteExerciseResp, error)
	ListExpense(context.Context, ListExerciseReq) (ListWorkoutResp, error)
}

type Workout interface {
	CreateWorkout(context.Context, CreateWorkoutReq) (CreateWorkoutResp, error)
	UpdateWorkout(context.Context, UpdateWorkoutReq) (UpdateWorkoutResp, error)
	DeleteWorkout(context.Context, DeleteWorkoutReq) (DeleteWorkoutResp, error)
	ListWorkout(context.Context, ListWorkoutReq) (ListWorkoutResp, error)
	GenerateWorkoutReports(context.Context, GenerateWorkoutReportsReq) (GenerateWorkoutReportsRes, error)
}

type Database struct {
	DB *sql.DB
}

func InitDB() (*Database, error) {
	connStr := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable" // Замени пароль

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Database is not responding: %v", err))
	}

	fmt.Println("Connected to the database")
	return &Database{DB: db}, nil
}
