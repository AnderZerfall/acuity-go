package infrastructure

import (
	"Acuity/internal/domain/emotion"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/nlpodyssey/cybertron/pkg/tasks"
	"github.com/nlpodyssey/cybertron/pkg/tasks/textclassification"
)

type Classifier struct {
	service textclassification.Interface
}

var (
	once     sync.Once
	instance *Classifier
)

func NewClassifier() *Classifier {
	once.Do(func() {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("failed to get working directory:", err)
		}

		modelsDir, err := filepath.Abs(filepath.Join(cwd, "external"))
		if err != nil {
			log.Fatal("failed to resolve models directory:", err)
		}

		s, err := tasks.Load[textclassification.Interface](&tasks.Config{
			ModelsDir:        modelsDir,
			ModelName:        "nlp-model",
			DownloadPolicy:   tasks.DownloadNever,
			ConversionPolicy: tasks.ConvertNever,
		})
		if err != nil {
			log.Fatal("failed to load model:", err)
		}

		instance = &Classifier{service: s}
	})
	return instance
}

func (c *Classifier) ClassifyEmotion(ctx context.Context, text string) (emotion.EmotionResult, error) {
	resp, err := c.service.Classify(ctx, text)
	if err != nil {
		return emotion.EmotionResult{}, fmt.Errorf("classify: %w", err)
	}

	return emotion.EmotionResult{
		Labels: resp.Labels,
		Scores: resp.Scores,
	}, nil
}
