package neurons


type Perceptron struct {
    FeaturesNumber     int
    IterationsNumber   int
    LearningStep       float64
    Weights          []float64
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

    p.Weights = make([]float64, featuresNumber + 1)
    p.Errors  = make([]int, 0)

    return p
}


func (p *Perceptron) Reset() {
    for i := range p.Weights {
        p.Weights[i] = 0
    }

    p.Errors = make([]int, p.IterationsNumber)
    for i := range p.Errors {
        p.Errors[i] = 0
    }
}


func (p *Perceptron) Retrain(trainingSet []TrainingSetRow) {
    p.Reset()

    for i := 0; i < p.IterationsNumber; i++ {
        errors := 0

        for _, row := range trainingSet {
            predicted := p.Predict(row.Features)
            dw0 := p.LearningStep * (row.Expected - predicted)

            p.Weights[0] += dw0
            for i := 0; i < p.FeaturesNumber; i++ {
                p.Weights[i + 1] += dw0 * row.Features[i]
            }

            if dw0 != 0 {
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
    result := p.Weights[0]

    for i := 0; i < p.FeaturesNumber; i++ {
        result += p.Weights[i + 1] * features[i]
    }

    return result
}
