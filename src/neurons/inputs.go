package neurons


type FeaturesRow []float64

type LabeledFeaturesRow struct {
    Features FeaturesRow
    Label    float64
}

type LabeledFeaturesSeries []LabeledFeaturesRow
