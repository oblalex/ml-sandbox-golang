package neurons


type FeaturesRow []float64

type TrainingSetRow struct {
    Features FeaturesRow
    Expected float64
}
