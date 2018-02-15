package main

import (
    "encoding/csv"
    "fmt"
    "net/http"
    "neurons"
    "strconv"
)


func readCSVFromURL(target string) ([][]string, error) {
	resp, err := http.Get(target)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func prepareTrainingSet(data [][]string) ([]neurons.TrainingSetRow, error) {
    result := make([]neurons.TrainingSetRow, 100)

    for i, row := range data[:100] {
        sepalLength, err := strconv.ParseFloat(row[0], 64)
        if err != nil {
            return nil, err
        }

        petalLength, err := strconv.ParseFloat(row[2], 64)
        if err != nil {
            return nil, err
        }

        var label float64

        if row[4] == "Iris-setosa" {
            label = -1.0
        } else {
            label = +1.0
        }

        result[i] = neurons.TrainingSetRow{
            neurons.FeaturesRow{sepalLength, petalLength},
            label,
        }
    }

    return result, nil
}


func main() {
    dataURL   := "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"

    data, err := readCSVFromURL(dataURL)
	if err != nil {
		panic(err)
	}

    trainingSet, err := prepareTrainingSet(data)
    if err != nil {
        panic(err)
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
        p.Weights(),
        p.Errors,
    )

    predicted := p.Predict(neurons.FeaturesRow{4.0, 2.0})
    fmt.Printf("predicted 1: %+1.0f, expected: +1\n", predicted)

    predicted = p.Predict(neurons.FeaturesRow{4.0, 1.0})
    fmt.Printf("predicted 2: %+1.0f, expected: -1\n", predicted)
}
