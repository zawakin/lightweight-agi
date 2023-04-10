package agi

import (
	"fmt"
	"strings"
)

// Objective is the objective of the task.
// It is used to determine the type of task to create.
type Objective string

func (o Objective) String() string {
	return string(o)
}

type Milestone struct {
	Objective Objective `json:"objective"`
	Name      string    `json:"name"`
}

type Milestones []Milestone

func (ms Milestones) String() string {
	var sb strings.Builder
	for i, m := range ms {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, m.Name))
	}
	return sb.String()
}
