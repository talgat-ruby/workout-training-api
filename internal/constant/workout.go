package constant

type WorkoutStatus string

const (
	StatusPending    WorkoutStatus = "pending"
	StatusInProgress WorkoutStatus = "in_progress"
	StatusCompleted  WorkoutStatus = "completed"
	StatusCancelled  WorkoutStatus = "cancelled"
)

// MuscleGroup represents different muscle groups
type MuscleGroup string

const (
	MuscleGroupChest      MuscleGroup = "chest"
	MuscleGroupBack       MuscleGroup = "back"
	MuscleGroupShoulders  MuscleGroup = "shoulders"
	MuscleGroupBiceps     MuscleGroup = "biceps"
	MuscleGroupTriceps    MuscleGroup = "triceps"
	MuscleGroupLegs       MuscleGroup = "legs"
	MuscleGroupCalves     MuscleGroup = "calves"
	MuscleGroupAbs        MuscleGroup = "abs"
	MuscleGroupForearms   MuscleGroup = "forearms"
	MuscleGroupTraps      MuscleGroup = "traps"
	MuscleGroupGlutes     MuscleGroup = "glutes"
	MuscleGroupHamstrings MuscleGroup = "hamstrings"
	MuscleGroupQuadriceps MuscleGroup = "quadriceps"
	MuscleGroupFullBody   MuscleGroup = "full_body"
	MuscleGroupCardio     MuscleGroup = "cardio"
)

// ExerciseCategory represents different types of exercises
type ExerciseCategory string

const (
	CategoryStrength    ExerciseCategory = "strength"
	CategoryCardio      ExerciseCategory = "cardio"
	CategoryFlexibility ExerciseCategory = "flexibility"
	CategoryBalance     ExerciseCategory = "balance"
)
