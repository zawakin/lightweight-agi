package agi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/examples/prompts"
	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
	"github.com/zawakin/lightweight-agi/prompt"
)

const (
	maxTaskIterationCount = 20
)

var (
	chunkSize = 200
)

type AGIAgent struct {
	completionClient llmclient.CompletionClient
	embeddingClient  llmclient.EmbeddingClient

	dataStore *datastore.DataStore
}

func NewAGIAgent(
	completionClient llmclient.CompletionClient,
	embeddingClient llmclient.EmbeddingClient,

	dataStore *datastore.DataStore,
) *AGIAgent {
	return &AGIAgent{
		completionClient: completionClient,
		embeddingClient:  embeddingClient,

		dataStore: dataStore,
	}
}

func (a *AGIAgent) RunAGIByObjective(ctx context.Context, objective prompts.Objective) error {
	runner := prompt.NewPromptRunner(a.completionClient)

	var objectiveRefinementOutput prompts.ObjectiveRefinementOutput
	err := runner.Run(ctx, prompts.ObjectRefinementPrompt, &prompts.ObjectiveRefinementInput{
		Objective: objective,
	}, &objectiveRefinementOutput)
	if err != nil {
		return err
	}

	var milestoneCreationOutput prompts.MilestoneCreationOutput
	err = runner.Run(ctx, prompts.MilestoneCreationPrompt, &prompts.MilestoneCreationInput{
		Objective: objective,
	}, &milestoneCreationOutput)
	if err != nil {
		return err
	}

	milestones := milestoneCreationOutput.Milestones

	for _, milestone := range milestones {
		err := a.RunAGIByMilestone(ctx, milestone)
		if err != nil {
			return err
		}
	}

	log.Printf("Finished executing tasks for objective %s", objective)
	return nil
}

func (a *AGIAgent) RunAGIByMilestone(ctx context.Context, milestone prompts.Milestone) error {
	runner := prompt.NewPromptRunner(a.completionClient)

	objective := milestone.Objective

	var objectiveRefinementOutput prompts.ObjectiveRefinementOutput
	err := runner.Run(ctx, prompts.ObjectRefinementPrompt, &prompts.ObjectiveRefinementInput{
		Objective: objective,
	}, &objectiveRefinementOutput)
	if err != nil {
		return err
	}

	objective = objectiveRefinementOutput.RefinedObjective

	var taskCreationOutput prompts.TaskCreationOutput
	err = runner.Run(ctx, prompts.TaskCreationPrompt, &prompts.TaskCreationInput{
		Objective: objective,
	}, &taskCreationOutput)
	if err != nil {
		return err
	}

	tasks := taskCreationOutput.Tasks

	var solvedTasks prompts.Tasks

	for i := 0; i < maxTaskIterationCount; i++ {
		if len(tasks) == 0 {
			log.Printf("No more tasks to execute for objective %s\n", objective)
			break
		}

		task, remain := tasks.PopLeft()
		taskName := task.Name

		tasks = remain

		queryResults, err := a.dataStore.Query(ctx, []model.Query{
			{
				Query: taskName,
				TopK:  5,
			},
		})
		if err != nil {
			return err
		}
		if len(queryResults) != 1 {
			return fmt.Errorf("unexpected chunks length: %d", len(queryResults))
		}
		queryResult := queryResults[0]

		var relevantContext []prompts.TaskContext
		for _, chunk := range queryResult.Results {
			relevantContext = append(relevantContext, prompts.TaskContext{
				Text: chunk.Text,
			})
		}

		var executionOutput prompts.ExecutionOutput
		err = runner.Run(ctx, prompts.ExecutionPrompt, &prompts.ExecutionInput{
			Objective:       objective,
			CurrentTask:     task,
			SolvedTasks:     solvedTasks,
			RelevantContext: relevantContext,
		}, &executionOutput)
		if err != nil {
			return err
		}

		result := executionOutput.CurrentTaskResult

		var evaluationTaskOutput prompts.EvaluationTaskOutput
		err = runner.Run(ctx, prompts.EvaluationTasksPrompt, &prompts.EvaluationTaskInput{
			Objective:  objective,
			Task:       task,
			TaskResult: result,
		}, &evaluationTaskOutput)
		if err != nil {
			return err
		}

		documentID := model.NewDocumentID()
		_, err = a.dataStore.Upsert(ctx, []model.Document{
			{
				ID:   documentID,
				Text: result.ResultText,
			},
		}, &chunkSize)
		if err != nil {
			return err
		}

		var prioritizationOutput prompts.PriorizationOutput
		err = runner.Run(ctx, prompts.PrioritizationPrompt, &prompts.PrioritizationInput{
			Objective: objective,
			Tasks:     tasks,
		}, &prioritizationOutput)
		if err != nil {
			return err
		}

		tasks = prioritizationOutput.Tasks

		time.Sleep(1 * time.Second)
	}

	log.Printf("Finished executing tasks for objective %s", objective)

	return nil
}
