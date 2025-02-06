package workout

import (
	"database/sql"
	"time"
	"workout-training-api/internal/constant"
)

type User struct {
	UserID         string     `db:"user_id" json:"user_id"`
	Email          string     `db:"email" json:"email"`
	HashedPassword string     `db:"hashed_password" json:"-"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
}

// Workout represents a workout session
type Workout struct {
	WorkoutID     string                 `db:"workout_id" json:"workout_id"`
	UserID        string                 `db:"user_id" json:"user_id"`
	Name          string                 `db:"name" json:"name"`
	Description   sql.NullString         `db:"description" json:"description,omitempty"`
	Exercises     []Exercise             `json:"exercises,omitempty"` // Loaded separately
	Status        constant.WorkoutStatus `db:"status" json:"status"`
	Comments      []Comment              `json:"comments,omitempty"` // Loaded separately
	ScheduledDate []*time.Time           `db:"scheduled_date" json:"scheduled_date,omitempty"`
	CreatedAt     *time.Time             `db:"created_at" json:"created_at"`
	UpdatedAt     *time.Time             `db:"updated_at" json:"updated_at"`
}

// Exercise represents a single exercise within a workout
type Exercise struct {
	ExerciseID  string                    `db:"exercise_id" json:"exercise_id"`
	WorkoutID   string                    `db:"workout_id" json:"workout_id"`
	MuscleGroup constant.MuscleGroup      `db:"muscle_group" json:"muscle_group"`
	Category    constant.ExerciseCategory `db:"category" json:"category"`
	Name        string                    `db:"name" json:"name"`
	Sets        int                       `db:"sets" json:"sets"`
	RepsPerSet  int                       `db:"reps_per_set" json:"reps_per_set"`
	WeightKg    float64                   `db:"weight_kg" json:"weight_kg,omitempty"`
	Notes       string                    `db:"notes" json:"notes,omitempty"`
	CreatedAt   *time.Time                `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time                `db:"updated_at" json:"updated_at"`
}

// Comment represents a comment on a workout
type Comment struct {
	CommentID string     `db:"comment_id" json:"comment_id"`
	WorkoutID string     `db:"workout_id" json:"workout_id"`
	Comment   string     `db:"comment" json:"comment"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
