package agi

import (
	"context"
	"log"
	"time"

	"github.com/zawakin/lightweight-agi/agi/agents/evaluationagent"
	"github.com/zawakin/lightweight-agi/agi/agents/executionagent"
	"github.com/zawakin/lightweight-agi/agi/agents/objectiverefinementagent"
	"github.com/zawakin/lightweight-agi/agi/agents/priorizationagent"
	"github.com/zawakin/lightweight-agi/agi/agents/taskcontextagent"
	"github.com/zawakin/lightweight-agi/agi/agents/taskcreationagent"
	"github.com/zawakin/lightweight-agi/agi/model"
)

const (
	maxTaskIterationCount = 20
)

type AGIAgent struct {
	objectiveRefinementAgent objectiverefinementagent.ObjectiveRefinementAgent
	executionAgent           executionagent.ExecutionAgent
	evaluationAgent          evaluationagent.EvaluationAgent
	taskCreationAgent        taskcreationagent.TaskCreationAgent
	priorizationAgent        priorizationagent.PriorizationAgent
	taskContextAgent         taskcontextagent.TaskContextAgent
}

func NewAGIAgent(
	objectiveRefinementAgent objectiverefinementagent.ObjectiveRefinementAgent,
	executionAgent executionagent.ExecutionAgent,
	evaluationAgent evaluationagent.EvaluationAgent,
	taskCreationAgent taskcreationagent.TaskCreationAgent,
	priorizationAgent priorizationagent.PriorizationAgent,
	taskContextAgent taskcontextagent.TaskContextAgent,
) *AGIAgent {
	return &AGIAgent{
		objectiveRefinementAgent: objectiveRefinementAgent,
		executionAgent:           executionAgent,
		evaluationAgent:          evaluationAgent,
		taskCreationAgent:        taskCreationAgent,
		priorizationAgent:        priorizationAgent,
		taskContextAgent:         taskContextAgent,
	}
}

func (a *AGIAgent) RunAGIByObjective(ctx context.Context, objective model.Objective) error {
	LogStep("Objective", objective)

	objective, err := a.objectiveRefinementAgent.RefineObjective(ctx, objective)
	if err != nil {
		return err
	}

	LogStep("Refined Objective", objective)

	milestones, err := a.taskCreationAgent.CreateMilestones(ctx, objective)
	if err != nil {
		return err
	}

	LogStep("Milestones", milestones)

	for _, milestone := range milestones {
		err := a.RunAGIByMilestone(ctx, milestone)
		if err != nil {
			return err
		}
	}

	log.Printf("Finished executing tasks for objective %s", objective)
	return nil
}

func (a *AGIAgent) RunAGIByMilestone(ctx context.Context, milestone model.Milestone) error {
	objective := milestone.Objective

	LogStep("Milestone Objective", objective)

	objective, err := a.objectiveRefinementAgent.RefineObjective(ctx, objective)
	if err != nil {
		return err
	}

	LogStep("Refined Milestone Objective", objective)

	tasks, err := a.taskCreationAgent.CreateTasks(ctx, objective, nil, nil)
	if err != nil {
		return err
	}

	for i := 0; i < maxTaskIterationCount; i++ {
		if len(tasks) == 0 {
			log.Printf("No more tasks to execute for objective %s\n", objective)
			break
		}

		task, remain := tasks.PopLeft()
		taskName := task.Name

		tasks = remain

		LogStep("Task", taskName)

		relevantContext, err := a.taskContextAgent.FindRelevantContext(ctx, task)
		if err != nil {
			return err
		}

		LogStep("Relevant Context", relevantContext)

		result, err := a.executionAgent.ExecuteTask(ctx, objective, task, relevantContext)
		if err != nil {
			return err
		}

		LogStep("Task Result", result.ResultText)

		evaluation, err := a.evaluationAgent.EvaluateTask(ctx, objective, task, result)
		if err != nil {
			return err
		}

		LogStep("Task Evaluation", evaluation)

		err = a.taskContextAgent.StoreContext(ctx, model.TaskContext{
			Task: task,
			Text: result.ResultText,
		})
		if err != nil {
			return err
		}

		prioritizedTasks, err := a.priorizationAgent.PrioritizeTasks(ctx, objective, tasks)
		if err != nil {
			return err
		}

		LogStep("Prioritized Tasks", prioritizedTasks)

		tasks = prioritizedTasks

		err = a.dumpAllSavedContext(ctx, objective, tasks)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("Finished executing tasks for objective %s", objective)

	return nil
}
