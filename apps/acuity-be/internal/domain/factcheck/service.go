package factcheck

import "context"

type FactResult struct {
	Score float64
}

type FactCheckService interface {
	CheckText(ctx context.Context, text string) (FactResult, error)
}
