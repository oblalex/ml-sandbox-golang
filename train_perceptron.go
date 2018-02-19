package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"
)

import "neurons"


const (
    DATA_DIR_PATH = "./data"
    TRAINING_SET_FILE_PATH = DATA_DIR_PATH + "/" + "training-set"
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


func main() {
    trainingSet, err := loadTrainingSet(TRAINING_SET_FILE_PATH)
    if err != nil {
        panic(err)
    }

    p := neurons.NewPerceptron(2, 10, 0.1)
    p.Train(trainingSet)

    fmt.Printf(
        "features number   = %d\n" +
        "iterations number = %d\n" +
        "learning step     = %f\n" +
        "weights           = %v\n" +
        "errors            = %v\n",
        p.FeaturesNumber,
        p.IterationsNumber,
        p.LearningStep,
        p.Weights(),
        p.Errors,
    )
}
