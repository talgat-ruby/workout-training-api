package workout

type Exercise struct {
	ID          string
	Name        string
	Description string
	Category    string
	Repetitions int
	Sets        int
	Weight      float64
}

func (e *Exercise) GetName() string {
	return e.Name
}

func (e *Exercise) GetExerciseIDs() string {
	return e.ID
}

func (e *Exercise) GetDescription() string {
	return e.Description
}

func (e *Exercise) GetCategory() string {
	return e.Category
}

func (e *Exercise) GetRepetitions() int {
	return e.Repetitions
}

func (e *Exercise) GetSets() int {
	return e.Sets
}

func (e *Exercise) GetWeight() float64 {
	return e.Weight
}
