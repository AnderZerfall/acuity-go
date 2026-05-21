package profile

import "context"

type ProfileaResult struct {
	Labels []string
	Scores []float64
}

type ProfileService interface {
	CheckProfile(ctx context.Context, text string) (ProfileaResult, error)
}
