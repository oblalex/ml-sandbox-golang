package main

import "fmt"
import "neurons"


func main() {
    trainingSet := []neurons.TrainingSetRow{
        neurons.TrainingSetRow{[]float64{5.1, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.9, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.7, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{4.6, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{5.4, 1.7}, -1},
        neurons.TrainingSetRow{[]float64{4.6, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{4.4, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.9, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.4, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{4.8, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{4.8, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.3, 1.1}, -1},
        neurons.TrainingSetRow{[]float64{5.8, 1.2}, -1},
        neurons.TrainingSetRow{[]float64{5.7, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.4, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{5.7, 1.7}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.4, 1.7}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{4.6, 1.0}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.7}, -1},
        neurons.TrainingSetRow{[]float64{4.8, 1.9}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{5.2, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.2, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.7, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{4.8, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{5.4, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.2, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.5, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{4.9, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.2}, -1},
        neurons.TrainingSetRow{[]float64{5.5, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{4.9, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{4.4, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{4.5, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{4.4, 1.3}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.9}, -1},
        neurons.TrainingSetRow{[]float64{4.8, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{5.1, 1.6}, -1},
        neurons.TrainingSetRow{[]float64{4.6, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{5.3, 1.5}, -1},
        neurons.TrainingSetRow{[]float64{5.0, 1.4}, -1},
        neurons.TrainingSetRow{[]float64{7.0, 4.7}, 1},
        neurons.TrainingSetRow{[]float64{6.4, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{6.9, 4.9}, 1},
        neurons.TrainingSetRow{[]float64{5.5, 4.0}, 1},
        neurons.TrainingSetRow{[]float64{6.5, 4.6}, 1},
        neurons.TrainingSetRow{[]float64{5.7, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{6.3, 4.7}, 1},
        neurons.TrainingSetRow{[]float64{4.9, 3.3}, 1},
        neurons.TrainingSetRow{[]float64{6.6, 4.6}, 1},
        neurons.TrainingSetRow{[]float64{5.2, 3.9}, 1},
        neurons.TrainingSetRow{[]float64{5.0, 3.5}, 1},
        neurons.TrainingSetRow{[]float64{5.9, 4.2}, 1},
        neurons.TrainingSetRow{[]float64{6.0, 4.0}, 1},
        neurons.TrainingSetRow{[]float64{6.1, 4.7}, 1},
        neurons.TrainingSetRow{[]float64{5.6, 3.6}, 1},
        neurons.TrainingSetRow{[]float64{6.7, 4.4}, 1},
        neurons.TrainingSetRow{[]float64{5.6, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{5.8, 4.1}, 1},
        neurons.TrainingSetRow{[]float64{6.2, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{5.6, 3.9}, 1},
        neurons.TrainingSetRow{[]float64{5.9, 4.8}, 1},
        neurons.TrainingSetRow{[]float64{6.1, 4.0}, 1},
        neurons.TrainingSetRow{[]float64{6.3, 4.9}, 1},
        neurons.TrainingSetRow{[]float64{6.1, 4.7}, 1},
        neurons.TrainingSetRow{[]float64{6.4, 4.3}, 1},
        neurons.TrainingSetRow{[]float64{6.6, 4.4}, 1},
        neurons.TrainingSetRow{[]float64{6.8, 4.8}, 1},
        neurons.TrainingSetRow{[]float64{6.7, 5.0}, 1},
        neurons.TrainingSetRow{[]float64{6.0, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{5.7, 3.5}, 1},
        neurons.TrainingSetRow{[]float64{5.5, 3.8}, 1},
        neurons.TrainingSetRow{[]float64{5.5, 3.7}, 1},
        neurons.TrainingSetRow{[]float64{5.8, 3.9}, 1},
        neurons.TrainingSetRow{[]float64{6.0, 5.1}, 1},
        neurons.TrainingSetRow{[]float64{5.4, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{6.0, 4.5}, 1},
        neurons.TrainingSetRow{[]float64{6.7, 4.7}, 1},
        neurons.TrainingSetRow{[]float64{6.3, 4.4}, 1},
        neurons.TrainingSetRow{[]float64{5.6, 4.1}, 1},
        neurons.TrainingSetRow{[]float64{5.5, 4.0}, 1},
        neurons.TrainingSetRow{[]float64{5.5, 4.4}, 1},
        neurons.TrainingSetRow{[]float64{6.1, 4.6}, 1},
        neurons.TrainingSetRow{[]float64{5.8, 4.0}, 1},
        neurons.TrainingSetRow{[]float64{5.0, 3.3}, 1},
        neurons.TrainingSetRow{[]float64{5.6, 4.2}, 1},
        neurons.TrainingSetRow{[]float64{5.7, 4.2}, 1},
        neurons.TrainingSetRow{[]float64{5.7, 4.2}, 1},
        neurons.TrainingSetRow{[]float64{6.2, 4.3}, 1},
        neurons.TrainingSetRow{[]float64{5.1, 3.0}, 1},
        neurons.TrainingSetRow{[]float64{5.7, 4.1}, 1},
    }
    p := neurons.NewPerceptron(2, 10, 0.1)

    p.Retrain(trainingSet)

    fmt.Printf(
        "features number   = %d\n" +
        "iterations number = %d\n" +
        "learning step     = %f\n" +
        "weights           = %v\n" +
        "errors            = %v\n",
        p.FeaturesNumber,
        p.IterationsNumber,
        p.LearningStep,
        p.Weights,
        p.Errors,
    )

    prediction := p.Predict([]float64{4.0, 2.0})
    fmt.Printf("prediction 1: %+1d, expected +1\n", prediction)

    prediction = p.Predict([]float64{4.0, 1.0})
    fmt.Printf("prediction 2: %+1d, expected -1\n", prediction)
}
