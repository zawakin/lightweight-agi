package evaluationagent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zawakin/lightweight-agi/agi/model"
)

func ParseTaskEvaluationFromString(s string) (model.TaskEvaluation, error) {
	lines := strings.Split(s, "\n")

	if len(lines) < 2 {
		return model.TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
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
				return model.TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
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
		return model.TaskEvaluation{}, fmt.Errorf("invalid task evaluation format: %s", s)
	}

	return model.TaskEvaluation{
		Score:  score,
		Reason: reason,
	}, nil
}
