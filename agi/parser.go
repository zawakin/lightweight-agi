package agi

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseMilestones(s string) ([]Milestone, error) {
	splitMilestones := strings.Split(s, "\n")
	var result []Milestone
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

		result = append(result, Milestone{
			Objective: Objective(strings.TrimSpace(indexName[1])),
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
func ParseTasksFromString(s string) (Tasks, error) {
	splitTasks := strings.Split(s, "\n")
	var result Tasks
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
			return Tasks{}, fmt.Errorf("invalid task format: %s", task)
		}

		result = append(result, Task{
			ID:   MakeTaskID(),
			Name: strings.TrimSpace(indexName[1]),
		})
	}

	return result, nil
}

func ParseTaskEvaluationFromString(s string) (TaskEvaluation, error) {
	lines := strings.Split(s, "\n")

	if len(lines) < 2 {
		return TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
	}

	var score int
	var reason string
	var err error

	hasScore := false

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Score:") {
			score, err = strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "Score:")))
			if err != nil {
				return TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
			}
			hasScore = true
			continue
		}
		if strings.HasPrefix(line, "Reason:") {
			reason = strings.TrimSpace(strings.TrimPrefix(line, "Reason:"))
			continue
		}
	}

	if !hasScore {
		return TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
	}

	return TaskEvaluation{
		Score:  score,
		Reason: reason,
	}, nil
}
