package emotion

import "context"

type EmotionResult struct {
	Labels []string
	Scores []float64
}

type EmotionClassifierService interface {
	ClassifyEmotion(ctx context.Context, text string) (EmotionResult, error)
}
