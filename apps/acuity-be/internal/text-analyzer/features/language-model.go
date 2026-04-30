package features

import (
	"context"
	"fmt"
	"log"

	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/textclassification"
)

func RunLanguageModel() {
	obj, err := tasks.Load[textclassification.Interface](&tasks.Config{
		ModelsDir: "../../external/nlp-model",
		ModelName: "../../external/nlp-model",
	})

	if err != nil {
		log.Fatal(err)
	}

	result, _ := obj.Classify(context.Background(), "Hey, analyze me!")

	fmt.Println(result.Labels, result.Scores)
}
