package prompts

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

var (
	// MaxContextLength is the maximum length of the context that can be
	// returned by the model.
	MaxContextLength = 2000
)

// Task is a struct that contains the base task information.
// It is used to identify the task and to provide a name for the task.
type Task struct {
	Name string `json:"name"`
}

func (t Task) String() string {
	return fmt.Sprintf(`"%s"`, t.Name)
}

type Tasks []Task

func (ts *Tasks) Add(task Task) {
	*ts = append(*ts, task)
}

func (ts Tasks) PopLeft() (Task, Tasks) {
	return ts[0], ts[1:]
}

func (ts Tasks) String() string {
	var ss []string
	for _, t := range ts {
		ss = append(ss, t.String())
	}
	return fmt.Sprintf(`[%s]`, strings.Join(ss, ","))
}

// SubTask represents a smaller unit of work within a Task.
type SubTask struct {
	Name string
}

type TaskResult struct {
	Task       Task   `json:"task"`
	ResultText string `json:"result_text"`
}

type TaskEvaluation struct {
	// Score has the range of 0 to 100.
	Score  int    `json:"score"`
	Reason string `json:"reason"`
}

func (t TaskEvaluation) String() string {
	return fmt.Sprintf(`%d%%: %s`, t.Score, t.Reason)
}

// TaskContext is a struct that contains the base task and the context
// for that task.
type TaskContext struct {
	Text string `json:"text"`
}
