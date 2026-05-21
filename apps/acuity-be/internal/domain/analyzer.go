package domain

import (
	"Acuity/gen/api/social_network/v1"
	"Acuity/internal/domain/emotion"
	"Acuity/internal/domain/factcheck"
	"context"
)

type PostAnalysisResult struct {
	TrustCoefficient float64
	Details          []string
}

type PostAnalyzer struct {
	emotionService     emotion.EmotionClassifierService
	factCheckerService factcheck.FactCheckService
	// mediaAnalyzer    MediaAnalyzer
}

const FallbackScore = 0.5

func NewPostAnalyzer(ec emotion.EmotionClassifierService, fc factcheck.FactCheckService) *PostAnalyzer {
	return &PostAnalyzer{emotionService: ec, factCheckerService: fc}
}

func (pa *PostAnalyzer) Analyze(ctx context.Context, post *social_network.Post) (PostAnalysisResult, error) {
	emotionResult := pa.getEmotionScore(ctx, post.Content)
	factcheckResult := pa.getFactcheckScore(ctx, post.Title)

	// 3. TODO: Call Media domain

	trustScore := emotionResult + factcheckResult

	return PostAnalysisResult{
		TrustCoefficient: trustScore,
		Details:          nil,
	}, nil
}

func (pa *PostAnalyzer) getEmotionScore(ctx context.Context, content string) float64 {
	emotionResult, err := pa.emotionService.ClassifyEmotion(ctx, content)

	if err != nil {
		return FallbackScore
	}

	weightedEmotions := 0.0

	for i := 0; i < len(emotionResult.Labels); i++ {
		key := emotion.Emotions(emotionResult.Labels[i])
		factor := emotion.EmotionWeights[key]

		weightedEmotions += factor * emotionResult.Scores[i]
	}

	return MethodsWeights[Emotions] * weightedEmotions
}

func (pa *PostAnalyzer) getFactcheckScore(ctx context.Context, title string) float64 {
	factResult, err := pa.factCheckerService.CheckText(ctx, title)

	if err != nil {
		return FallbackScore
	}

	return MethodsWeights[Factcheck] * factResult.Score
}
