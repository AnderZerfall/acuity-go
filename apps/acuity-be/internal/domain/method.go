package domain

type Methods string

const (
	Emotions  Methods = "emotions"
	Factcheck Methods = "fatchcheck"
	Media     Methods = "media"
	Proile    Methods = "profile"
)

// MethodsWeights shows how each emotion affects the trust score. 1 - no effect
var MethodsWeights = map[Methods]float64{
	Emotions:  1,
	Factcheck: 0,
	Media:     0,
	Proile:    0,
}
