package agi

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var (
	// MaxContextLength is the maximum length of the context that can be
	// returned by the model.
	MaxContextLength = 2000
)

type TaskID uuid.UUID

func MakeTaskID() TaskID {
	return TaskID(uuid.New())
}

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}

func (id TaskID) MarshalText() ([]byte, error) {
	return []byte(uuid.UUID(id).String()), nil
}

func ParseTaskIDFromString(s string) (TaskID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return TaskID{}, err
	}
	return TaskID(id), nil
}

// Task is a struct that contains the base task information.
// It is used to identify the task and to provide a name for the task.
type Task struct {
	ID   TaskID `json:"id"`
	Name string `json:"name"`
}

func (t Task) String() string {
	return fmt.Sprintf(`"%s"`, t.Name)
}

type Tasks []Task

func (ts Tasks) Add(task Task) {
	ts = append(ts, task)
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
	ID   string
	Name string
}

type TaskResult struct {
	Task       Task
	ResultText string
}

type TaskEvaluation struct {
	Score  int // 0-100
	Reason string
}

func (t TaskEvaluation) String() string {
	return fmt.Sprintf(`%d%%: %s`, t.Score, t.Reason)
}

// TaskContext is a struct that contains the base task and the context
// for that task.
type TaskContext struct {
	Task Task   `json:"task"`
	Text string `json:"text"`
}

// TaskRelevantContext is a struct that contains the base task and the
// relevant context for that task.
type TaskRelevantContext struct {
	BaseTask        Task
	RelevantContext []TaskContext
}

func (t TaskRelevantContext) String() string {
	var ss []string
	for _, c := range t.RelevantContext {
		ss = append(ss, fmt.Sprintf(c.Text))
	}
	result := strings.Join(ss, ",")

	if len(result) > MaxContextLength {
		result = result[:MaxContextLength]
	}
	return result
}

// TaskContextCollection is a struct that contains the base task, the
// relevant context for that task, and the current tasks.
type TaskContextCollection struct {
	Objective    Objective     `json:"objective"`
	Contexts     []TaskContext `json:"contexts"`
	CurrentTasks Tasks         `json:"current_tasks"`
}

func NewTaskContextCollection(objective Objective, contexts []TaskContext, currentTasks Tasks) TaskContextCollection {
	return TaskContextCollection{
		Objective:    objective,
		Contexts:     contexts,
		CurrentTasks: currentTasks,
	}
}

func GetTaskContentFromMetadata(metadata map[string]interface{}) (TaskContext, error) {
	id, ok := hasMetadata(metadata, "id")
	if !ok {
		return TaskContext{}, fmt.Errorf("metadata %v does not have id", metadata)
	}

	task, ok := hasMetadata(metadata, "task")
	if !ok {
		return TaskContext{}, fmt.Errorf("metadata %v does not have task", metadata)
	}

	content, ok := hasMetadata(metadata, "content")
	if !ok {
		return TaskContext{}, fmt.Errorf("metadata %v does not have content", metadata)
	}

	idParsed, err := ParseTaskIDFromString(id)
	if err != nil {
		return TaskContext{}, err
	}

	return TaskContext{
		Task: Task{
			ID:   idParsed,
			Name: task,
		},
		Text: content,
	}, nil
}

func hasMetadata(metadata map[string]interface{}, key string) (string, bool) {
	v, ok := metadata[key]
	if !ok {
		return "", false
	}

	s, ok := v.(string)
	if !ok {
		return "", false
	}

	return s, true
}
