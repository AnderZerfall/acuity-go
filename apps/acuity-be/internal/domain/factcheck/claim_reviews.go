package factcheck

import "strings"

type ClaimsReview string

const (
	False       ClaimsReview = "false"
	MostlyFalse ClaimsReview = "mostly false"
	PartlyFalse ClaimsReview = "partly false"
	HalfTrue    ClaimsReview = "half true"
	MostlyTrue  ClaimsReview = "mostly true"
	True        ClaimsReview = "true"
	Unknown     ClaimsReview = ""
)

// ClaimsReviewWeights shows how each claims affects the trust score.
// 0 - negative effect, 1 - positive effect
var ClaimsReviewWeights = map[ClaimsReview]float64{
	Unknown:     0.5,
	False:       0,
	MostlyFalse: 0.2,
	PartlyFalse: 0.4,
	HalfTrue:    0.6,
	MostlyTrue:  0.8,
	True:        1,
}

func MapToEnum(input string) ClaimsReview {
	cleaned := strings.ToLower(strings.TrimSpace(input))

	if cleaned == "" {
		return Unknown
	}

	switch {
	case strings.Contains(cleaned, "mostly false") || strings.Contains(cleaned, "mostly untrue"):
		return MostlyFalse
	case strings.Contains(cleaned, "partly false") || strings.Contains(cleaned, "partially false"):
		return PartlyFalse
	case strings.Contains(cleaned, "half true") || strings.Contains(cleaned, "half false"):
		return HalfTrue
	case strings.Contains(cleaned, "mostly true"):
		return MostlyTrue
	case strings.Contains(cleaned, "false") || strings.Contains(cleaned, "fake") || strings.Contains(cleaned, "untrue"):
		return False
	case strings.Contains(cleaned, "true") || strings.Contains(cleaned, "correct"):
		return True
	default:
		return Unknown
	}
}
