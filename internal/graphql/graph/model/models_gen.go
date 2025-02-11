// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Exercise struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Repetitions int32   `json:"repetitions"`
	Sets        int32   `json:"sets"`
	Weight      float64 `json:"weight"`
}

type ExerciseInput struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Repetitions int32   `json:"repetitions"`
	Sets        int32   `json:"sets"`
	Weight      float64 `json:"weight"`
}

type Mutation struct {
}

type Ping struct {
	Message *string `json:"message,omitempty"`
}

type Query struct {
}

type SignInResp struct {
	Token string `json:"token"`
}

type SignUpResp struct {
	Token string `json:"token"`
}

type Workout struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Exercises      []*Exercise `json:"exercises"`
	Description    string      `json:"description"`
	ScheduledTimes int32       `json:"scheduledTimes"`
}

type WorkoutInput struct {
	ID             string           `json:"id"`
	Name           string           `json:"name"`
	Exercises      []*ExerciseInput `json:"exercises"`
	Description    string           `json:"description"`
	ScheduledTimes int32            `json:"scheduledTimes"`
}
