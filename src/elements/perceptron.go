package elements

type Perceptron struct {
    FeaturesNumber     int
    IterationsNumber   int
    LearningStep       float64

    weights         []float64
    errors          []float64
}

func NewPerceptron(
    featuresNumber   int,
    iterationsNumber int,
    learningStep     float64,
) *Perceptron {
    p := new(Perceptron)

    p.FeaturesNumber   = featuresNumber
    p.IterationsNumber = iterationsNumber
    p.LearningStep     = learningStep

    p.weights = make([]float64, featuresNumber + 1)
    p.errors  = make([]float64, iterationsNumber)

    return p
}
