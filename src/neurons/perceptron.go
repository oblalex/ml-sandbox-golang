package neurons


type Perceptron struct {
    FeaturesNumber     int
    IterationsNumber   int
    LearningStep       float64
    weight0            float64
    weightFeatures   []float64
    Errors           []int
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
    p.weightFeatures   = make([]float64, featuresNumber)

    p.Reset()
    return p
}


func (p *Perceptron) Weights() []float64 {
    result := make([]float64, p.FeaturesNumber + 1)
    result[0] = p.weight0
    copy(result[1:], p.weightFeatures[:])
    return result
}


func (p *Perceptron) Reset() {
    p.weight0 = 0
    for i := range p.weightFeatures {
        p.weightFeatures[i] = 0
    }

    if p.Errors == nil {
        p.Errors = make([]int, p.IterationsNumber)
    }
    for i := range p.Errors {
        p.Errors[i] = 0
    }
}


func (p *Perceptron) Train(trainingSet LabeledFeaturesSeries) {
    for i := 0; i < p.IterationsNumber; i++ {
        errors := 0

        for _, row := range trainingSet {
            predicted := p.Predict(row.Features)
            cost := p.LearningStep * (row.Label - predicted)

            p.weight0 += cost

            for i := 0; i < p.FeaturesNumber; i++ {
                p.weightFeatures[i] += cost * row.Features[i]
            }

            if cost != 0 {
                errors += 1
            }
        }

        p.Errors[i] = errors
    }
}


func (p *Perceptron) Predict(features FeaturesRow) float64 {
    netInput := p.NetInput(features)

    if netInput >= 0 {
        return 1
    } else {
        return -1
    }
}


func (p *Perceptron) NetInput(features FeaturesRow) float64 {
    result := p.weight0

    for i := 0; i < p.FeaturesNumber; i++ {
        result += p.weightFeatures[i] * features[i]
    }

    return result
}
