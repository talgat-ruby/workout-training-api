package constant

type ExpenseSortByField string

const (
	ExpenseSortByFieldUnspecified ExpenseSortByField = "unspecified"
	ExpenseSortByFieldID          ExpenseSortByField = "id"
	ExpenseSortByFieldAmount      ExpenseSortByField = "amount"
	ExpenseSortByFieldCategory    ExpenseSortByField = "category"
	ExpenseSortByFieldCreatedAt   ExpenseSortByField = "createdAt"
	ExpenseSortByFieldUpdatedAt   ExpenseSortByField = "updatedAt"
)

type Status int

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusInProgress
	StatusCompleted
	StatusCancelled
)

type MuscleGroup int

const (
	MuscleGroupUnspecified MuscleGroup = iota
	MuscleGroupChest
	MuscleGroupBack
	MuscleGroupShoulders
	MuscleGroupBiceps
	MuscleGroupTriceps
	MuscleGroupLegs
	MuscleGroupCalves
	MuscleGroupAbs
	MuscleGroupForearms
	MuscleGroupTraps
	MuscleGroupGlutes
	MuscleGroupHamstrings
	MuscleGroupQuadriceps
	MuscleGroupFullBody
	MuscleGroupCardio
)

type Category int

const (
	CategoryUnspecified Category = iota
	CategoryStrength
	CategoryCardio
	CategoryFlexibility
	CategoryBalance
)

/*



type Workout struct {
   ID            string      `json:"id"`
   UserID        string      `json:"user_id"`
   Name          string      `json:"name"`
   Description   string      `json:"description"`
   Exercises     []Exercise  `json:"exercises"`
   Status        Status      `json:"status"`
   Comments      []Comment   `json:"comments"`
   ScheduledDate time.Time   `json:"scheduled_date"`
   CreatedAt     time.Time   `json:"created_at"`
   UpdatedAt     time.Time   `json:"updated_at"`
}

type Exercise struct {
   ID          string      `json:"id"`
   MuscleGroup MuscleGroup `json:"muscle_group"`
   Category    Category    `json:"category"`
   Name        string      `json:"name"`
   Sets        uint32      `json:"sets"`
   RepsPerSet  uint32      `json:"reps_per_set"`
   WeightKg    float32     `json:"weight_kg"`
   Notes       string      `json:"notes"`
   CreatedAt   time.Time   `json:"created_at"`
   UpdatedAt   time.Time   `json:"updated_at"`
}

type Comment struct {
   Comment string `json:"comment"`
}

type Status int

const (
   StatusUnspecified Status = iota
   StatusPending
   StatusInProgress
   StatusCompleted
   StatusCancelled
)

type MuscleGroup int

const (
   MuscleGroupUnspecified MuscleGroup = iota
   MuscleGroupChest
   MuscleGroupBack
   MuscleGroupShoulders
   MuscleGroupBiceps
   MuscleGroupTriceps
   MuscleGroupLegs
   MuscleGroupCalves
   MuscleGroupAbs
   MuscleGroupForearms
   MuscleGroupTraps
   MuscleGroupGlutes
   MuscleGroupHamstrings
   MuscleGroupQuadriceps
   MuscleGroupFullBody
   MuscleGroupCardio
)

type Category int

const (
   CategoryUnspecified Category = iota
   CategoryStrength
   CategoryCardio
   CategoryFlexibility
   CategoryBalance
)
*/
