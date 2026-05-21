package media

import "context"

type MediaResult struct {
	Labels []string
	Scores []float64
}

// TO DO: Extend for more media types
type MediaService interface {
	CheckImage(ctx context.Context, text string) (MediaResult, error)
}
