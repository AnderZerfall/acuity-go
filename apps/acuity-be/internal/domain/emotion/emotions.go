package emotion

type Emotions string

const (
	Neutral   Emotions = "neutral"
	Anger     Emotions = "anger"
	Sadness   Emotions = "sadness"
	Love      Emotions = "love"
	Surprise  Emotions = "surprise"
	Fear      Emotions = "fear"
	Happiness Emotions = "happiness"
	Disgust   Emotions = "disgust"
	Shame     Emotions = "shame"
	Guilt     Emotions = "guilt"
	Confusion Emotions = "confusion"
	Desire    Emotions = "desire"
	Sarcasm   Emotions = "sarcasm"
)

// EmotionWeights shows how each emotion affects the trust score. 1 - no effect
var EmotionWeights = map[Emotions]float64{
	Neutral:   1,
	Anger:     0.5,
	Sadness:   0.6,
	Love:      0.9,
	Surprise:  1,
	Fear:      0.5,
	Happiness: 0.8,
	Disgust:   0.5,
	Shame:     0.4,
	Guilt:     0.5,
	Confusion: 0.8,
	Desire:    0.5,
	Sarcasm:   0.3,
}
