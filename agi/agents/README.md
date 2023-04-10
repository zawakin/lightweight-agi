# `agents/`

This directory contains the AGI agents.

## Agents

* Evaluation Agent: Evaluates the results of tasks.
* Execution Agent: Executes tasks.
* Objective Refinement Agent: Refines the main objective into smaller, more manageable objectives.
* Prioritization Agent: Prioritizes tasks based on their relevance and importance.
* Task Context Agent: Stores the context of tasks for future reference.
* Task Creation Agent: Generates tasks and milestones based on the refined objectives.

## `agents/xxxagent/`

* `agent.go` - The file for the interface and implementation of the agent.
* `prompt.go` - The file for the prompt of the agent.
* `parser.go` - The file for the parser of the agent.
