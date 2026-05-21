package infrastructure

import (
	"Acuity/internal/domain/factcheck"
	"context"
	"fmt"
	"log"
	"os"

	factchecktools "google.golang.org/api/factchecktools/v1alpha1"
	"google.golang.org/api/option"
)

type GoogleService struct {
	client *factchecktools.Service
}

func NewGoogleService(ctx context.Context) *GoogleService {
	apiKey := os.Getenv("GOOGLE_API_KEY")

	if apiKey == "" {
		log.Fatal("Google Api key is missing")
	}

	service, err := factchecktools.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		log.Fatalf("Failed to initialize Google Service: %v", err)
	}

	return &GoogleService{client: service}
}

func (s *GoogleService) Search(ctx context.Context, query string) (*factchecktools.GoogleFactcheckingFactchecktoolsV1alpha1FactCheckedClaimSearchResponse, error) {

	response, err := s.client.Claims.Search().
		Query(query).
		Context(ctx).
		Do()

	if err != nil {
		return nil, fmt.Errorf("google api call failed: %w", err)
	}

	return response, nil
}

func (s *GoogleService) ImageSearch(ctx context.Context, imageURI string) (*factchecktools.GoogleFactcheckingFactchecktoolsV1alpha1FactCheckedClaimImageSearchResponse, error) {
	response, err := s.client.Claims.ImageSearch().
		ImageUri(imageURI).
		Context(ctx).
		Do()

	if err != nil {
		return nil, fmt.Errorf("google api call failed: %w", err)
	}

	return response, nil
}

func (s *GoogleService) CheckText(ctx context.Context, text string) (factcheck.FactResult, error) {
	response, err := s.Search(ctx, text)

	if err != nil {
		return factcheck.FactResult{Score: 0.5}, err
	}

	if response == nil || len(response.Claims) == 0 {
		return factcheck.FactResult{
			Score: 0.5,
		}, nil
	}

	score := 0.0
	reviewsAmount := 0.0

	for _, claim := range response.Claims {
		for _, review := range claim.ClaimReview {
			fmt.Printf("%s", review.TextualRating)
			score += factcheck.ClaimsReviewWeights[factcheck.MapToEnum(review.TextualRating)]
			reviewsAmount += 1
		}
	}

	return factcheck.FactResult{
		Score: score / reviewsAmount,
	}, err
}

// func (s *GoogleService) CheckImage(ctx context.Context, text string) (factcheck.FactResult, error) {
// 	response, err := s.ImageSearch(ctx, text)

// 	// Add data transformation

// 	return response, err
// }
