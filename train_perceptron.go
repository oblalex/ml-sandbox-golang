package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
)

import "neurons"


const (
    DATA_DIR_PATH = "./data"
    TRAINING_SET_FILE_PATH = DATA_DIR_PATH + "/" + "training-set"
    PREDICTIONS_FILE_PATH  = DATA_DIR_PATH + "/" + "perceptron-predictions"
)


func loadTrainingSet(filePath string) (neurons.LabeledFeaturesSeries, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }

    var result neurons.LabeledFeaturesSeries
    reader := csv.NewReader(bufio.NewReader(file))

    for {
        row, error := reader.Read()

        if error == io.EOF {
            break
        } else if error != nil {
            return nil, err
        }

        sepalLength, err := strconv.ParseFloat(row[0], 64)
        if err != nil {
            return nil, err
        }

        petalLength, err := strconv.ParseFloat(row[1], 64)
        if err != nil {
            return nil, err
        }

        label, err := strconv.ParseFloat(row[2], 64)
        if err != nil {
            return nil, err
        }

        result = append(result, neurons.LabeledFeaturesRow{
            neurons.FeaturesRow{sepalLength, petalLength},
            label,
        })
    }

    return result, nil
}


func getTrainingSetBoundaries(
    trainingSet neurons.LabeledFeaturesSeries,
) (
    minSepalLength, maxSepalLength float64,
    minPetalLength, maxPetalLength float64,
) {
    maxSepalLength, maxPetalLength = 0.0, 0.0

    for _, row := range trainingSet {
        if row.Features[0] > maxSepalLength {
            maxSepalLength = row.Features[0]
        }
        if row.Features[1] > maxPetalLength {
            maxPetalLength = row.Features[1]
        }
    }

    minSepalLength, minPetalLength = maxSepalLength, maxPetalLength

    for _, row := range trainingSet {
        if row.Features[0] < minSepalLength {
            minSepalLength = row.Features[0]
        }
        if row.Features[1] < minPetalLength {
            minPetalLength = row.Features[1]
        }
    }

    return minSepalLength, maxSepalLength, minPetalLength, maxPetalLength
}


func generatePredictions(
    p *neurons.Perceptron,
    minSepalLength, maxSepalLength float64,
    minPetalLength, maxPetalLength float64,
    step float64,
) (neurons.LabeledFeaturesSeries) {

    sepalPoints := int(math.Ceil((maxSepalLength + step - minSepalLength) / step))
    petalPoints := int(math.Ceil((maxPetalLength + step - minPetalLength) / step))

    totalPoints := sepalPoints * petalPoints
    result      := make(neurons.LabeledFeaturesSeries, totalPoints)

    sepalLength := minSepalLength
    for i := 0; i < sepalPoints; i++ {

        offset      := i * petalPoints
        petalLength := minPetalLength
        for j := 0; j < petalPoints; j++ {

            features := neurons.FeaturesRow{sepalLength, petalLength}
            label    := p.Predict(features)
            idx      := offset + j

            result[idx]  = neurons.LabeledFeaturesRow{features, label}
            petalLength += step
        }

        sepalLength += step
    }

    return result
}


func main() {
    trainingSet, err := loadTrainingSet(TRAINING_SET_FILE_PATH)
    if err != nil {
        panic(err)
    }

    minSepalLength, maxSepalLength, minPetalLength, maxPetalLength := getTrainingSetBoundaries(
        trainingSet,
    )

    p := neurons.NewPerceptron(2, 10, 0.1)
    p.Train(trainingSet)

    fmt.Printf(
        "features number   = %d\n" +
        "iterations number = %d\n" +
        "learning step     = %f\n" +
        "weights           = %v\n" +
        "errors            = %v\n" +
        "min sepal length  = %f\n" +
        "max sepal length  = %f\n" +
        "min petal length  = %f\n" +
        "max petal length  = %f\n",
        p.FeaturesNumber,
        p.IterationsNumber,
        p.LearningStep,
        p.Weights(),
        p.Errors,
        minSepalLength,
        maxSepalLength,
        minPetalLength,
        maxPetalLength,
    )

    predictions := generatePredictions(
        p,
        minSepalLength - 1, maxSepalLength + 1,
        minPetalLength - 1, maxPetalLength + 1,
        0.1,
    )
    err = neurons.StoreLabeledFeaturesSeries(predictions, PREDICTIONS_FILE_PATH)
    if err != nil {
        panic(err)
    }
}
