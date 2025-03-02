package database

import (
	"context"
)

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUserByEmail(context.Context, FindUserReq) (FindUserResp, error)
}

//type CreateWorkoutReq interface {
//	GetName() string
//	GetDescription() string
//	GetExercises() []workout.Exercise
//	GetStatus() constant.WorkoutStatus
//	GetDate() time.Time
//	GetScheduledDate() []time.Time
//}

type Exercise interface {
	CreateExercise(context.Context, CreateExerciseReq) (CreateExerciseResp, error)
	UpdateExercise(context.Context, UpdateExerciseReq) (UpdateExerciseResp, error)
	DeleteExercise(context.Context, DeleteExerciseReq) (DeleteExerciseResp, error)
	ListExercise(context.Context, ListExerciseReq) (ListExerciseResp, error)
}

type Workout interface {
	CreateWorkout(context.Context, CreateWorkoutReq) (CreateWorkoutResp, error)
	UpdateWorkout(context.Context, UpdateWorkoutReq) (UpdateWorkoutResp, error)
	DeleteWorkout(context.Context, DeleteWorkoutReq) (DeleteWorkoutResp, error)
	ListWorkout(context.Context, ListWorkoutReq) (ListWorkoutResp, error)
	GenerateWorkoutReports(context.Context, GenerateWorkoutReportsReq) (GenerateWorkoutReportsRes, error)
}

type Database interface {
	User
	Exercise
	Workout
}
