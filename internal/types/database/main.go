package database

import "context"

type User interface {
	CreateUser(context.Context, CreateUserReq) (CreateUserResp, error)
	FindUser(ctx context.Context, req FindUserReq) (FindUserResp, error)
}

type Exercise interface{}

type Workout interface {
	CreateWorkout(context.Context, CreateWorkoutReq) (CreateWorkoutResp, error)
	UpdateWorkout(context.Context, UpdateWorkoutReq) (UpdateWorkoutResp, error)
	DeleteWorkout(context.Context, DeleteWorkoutReq) (DeleteWorkoutResp, error)
	ListWorkout(context.Context, ListWorkoutReq) (ListWorkoutResp, error)
}
