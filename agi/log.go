package agi

import (
	"context"
	"encoding/json"
	"log"
)

func (a *AGIAgent) dumpAllSavedContext(ctx context.Context, objective Objective, tasks Tasks) error {
	dumped, err := a.taskContextAgent.DebugDumpTaskContext(ctx)
	if err != nil {
		return err
	}

	collection := NewTaskContextCollection(objective, dumped, tasks)

	bs, err := json.Marshal(collection)
	if err != nil {
		return err
	}

	LogStep("[DEBUG] Dump All Saved Context", string(bs))

	return nil
}

func LogStep(step string, vs ...interface{}) {
	log.Println()
	log.Println("======= " + step + " ======")
	for _, v := range vs {
		log.Println(v)
	}
	log.Println()
}
