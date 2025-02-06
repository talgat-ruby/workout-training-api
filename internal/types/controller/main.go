package controller

import "context"

type Interceptor interface {
	Authenticator(context.Context, string) (context.Context, error)
}

type Auth interface {
	SignUp(context.Context, SignUpReq) (SignUpResp, error)
	SignIn(context.Context, SignInReq) (SignInResp, error)
}

type WorkoutCrud interface {
	CreateWorkout(context.Context, CreateWorkoutReq) (CreateWorkoutResp, error)
	UpdateWorkout(context.Context, UpdateWorkoutReq) (UpdateWorkoutResp, error)
	DeleteWorkout(context.Context, DeleteWorkoutReq) (DeleteWorkoutResp, error)
	ListWorkouts(context.Context, ListWorkoutsReq) (ListWorkoutsResp, error)
	ScheduleWorkout(context.Context, ScheduleWorkoutReq) (ScheduleWorkoutResp, error)
	GenerateReport(context.Context, GenerateReportReq) (GenerateReportResp, error)
}

type Controller interface {
	Interceptor
	Auth
	Workout
}
