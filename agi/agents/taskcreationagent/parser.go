package taskcreationagent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zawakin/lightweight-agi/agi/model"
)

func ParseMilestones(s string) ([]model.Milestone, error) {
	splitMilestones := strings.Split(s, "\n")
	var result []model.Milestone
	for _, milestone := range splitMilestones {
		if milestone == "" {
			continue
		}
		milestone = strings.TrimSpace(milestone)
		if milestone[0] < '0' || milestone[0] > '9' {
			continue
		}
		indexName := strings.SplitN(milestone, ".", 2)
		if len(indexName) != 2 {
			continue
		}
		if indexName[0] == "" || indexName[1] == "" {
			continue
		}
		_, err := strconv.Atoi(strings.TrimSpace(indexName[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid milestone format: %s", milestone)
		}

		result = append(result, model.Milestone{
			Objective: model.Objective(strings.TrimSpace(indexName[1])),
			Name:      strings.TrimSpace(indexName[1]),
		})
	}
	return result, nil
}

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
