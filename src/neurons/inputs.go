package neurons


type FeaturesRow []float64

type TrainingSetRow struct {
    Features FeaturesRow
    Label    float64
}
