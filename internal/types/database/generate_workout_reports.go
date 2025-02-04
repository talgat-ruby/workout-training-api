package database

type GenerateWorkoutReportsReq interface {
	GetUserID() string
}

type GenerateWorkoutReportsRes interface {
	GetProgress() string
	GetWorkouts() WorkoutListResp
}
