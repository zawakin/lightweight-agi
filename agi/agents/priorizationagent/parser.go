package priorizationagent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zawakin/lightweight-agi/agi/model"
)

// ParseTasksFromString parses a string of tasks into a Tasks object.
// The string should be in the following format:
// 1. {task1}
// 2. {task2}
// 3. {task3}
// ...
// The index of the task is ignored.
func ParseTasksFromString(s string) (model.Tasks, error) {
	splitTasks := strings.Split(s, "\n")
	var result model.Tasks
	for _, task := range splitTasks {
		if task == "" {
			continue
		}
		task = strings.TrimSpace(task)
		if task[0] < '0' || task[0] > '9' {
			continue
		}
		indexName := strings.SplitN(task, ".", 2)
		if len(indexName) != 2 {
			continue
		}
		if indexName[0] == "" || indexName[1] == "" {
			continue
		}
		_, err := strconv.Atoi(strings.TrimSpace(indexName[0]))
		if err != nil {
			return model.Tasks{}, fmt.Errorf("invalid task format: %s", task)
		}

		result = append(result, model.Task{
			ID:   model.MakeTaskID(),
			Name: strings.TrimSpace(indexName[1]),
		})
	}

	return result, nil
}
