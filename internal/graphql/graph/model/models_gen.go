package model

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type SignInResp struct {
	Token string `json:"token"`
}

type SignUpResp struct {
	Token string `json:"token"`
}

type Workout struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Exercises      []*Exercise `json:"exercises"`
	ScheduledTimes []time.Time `json:"date"`
}

// Exercise Copied and pasted from controller
type Exercise struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Repetitions int    `json:"repetitions"`
	Sets        int    `json:"sets"`
	Weight      int    `json:"weight"`
}

func (w *Workout) UnmarshalGQL(v any) error {
	data, ok := v.(map[string]any)
	if !ok {
		return fmt.Errorf("Workout should be an object")
	}

	if id, ok := data["id"].(string); ok {
		w.ID = id
	} else {
		return fmt.Errorf("Workout ID must be a string")
	}

	if name, ok := data["name"].(string); ok {
		w.Name = name
	} else {
		return fmt.Errorf("Workout Name must be a string")
	}

	return nil
}

func (w Workout) MarshalGQL(writer io.Writer) {
	jsonData, err := json.Marshal(w)
	if err != nil {
		fmt.Fprint(writer, "null")
		return
	}
	writer.Write(jsonData)
}
