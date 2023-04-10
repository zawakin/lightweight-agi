package agi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/zawakin/lightweight-agi/agi/model"
)

func (a *AGIAgent) dumpAllSavedContext(ctx context.Context, objective model.Objective, tasks model.Tasks) error {
	dumped, err := a.taskContextAgent.DebugDumpTaskContext(ctx)
	if err != nil {
		return err
	}

	collection := model.NewTaskContextCollection(objective, dumped, tasks)

	bs, err := json.Marshal(collection)
	if err != nil {
		return err
	}

	LogStep("[DEBUG] Dump All Saved Context", string(bs))

	return nil
}

// LogStep logs a step in the AGI process.
func LogStep(step string, vs ...interface{}) {
	log.Println()
	log.Println("======= " + step + " ======")
	for _, v := range vs {
		log.Println(v)
	}
	log.Println()
}
