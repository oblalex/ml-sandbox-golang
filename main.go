package main

import (
    "bytes"
    "crypto/sha1"
    "encoding/csv"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "path/filepath"
    "os"
    "strconv"
    "strings"
)

import "neurons"


func getStringHash(s string) (string) {
    h := sha1.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}


func getURLHash(url string) (string) {
    s := strings.SplitN(url, "://", 2)
    return getStringHash(s[1])
}


func downloadData(url string) ([]byte, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func storeData(filePath string, data []byte) (error) {
    dirPath := filepath.Dir(filePath)
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {

        err := os.Mkdir(dirPath, 0755)
        if err != nil {
            return err
        }
    }

    err := ioutil.WriteFile(filePath, data, 0644)
    if err != nil {
        return err
    }

    return nil
}


func maybeDownload(url string) ([]byte, error) {
    cacheDir := ".cache"
    fileName := getURLHash(url)
    filePath := fmt.Sprintf("%s/%s", cacheDir, fileName)

    data, err := ioutil.ReadFile(filePath)
    if err != nil {

        data, err = downloadData(url)
        if err != nil {
            return nil, err
        }

        err = storeData(filePath, data)
        if err != nil {
            return nil, err
        }
    }

    return data, nil
}


func readCSVdata(reader io.Reader) ([][]string, error) {
    csvReader := csv.NewReader(reader)

    data, err := csvReader.ReadAll()
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


func getTrainingSet(url string) ([]neurons.TrainingSetRow, error) {
    raw, err := maybeDownload(url)
    if err != nil {
        return nil, err
    }

    strings, err := readCSVdata(bytes.NewReader(raw))
    if err != nil {
        return nil, err
    }

    return prepareTrainingSet(strings)
}


func main() {
    dataUrl := "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"

    trainingSet, err := getTrainingSet(dataUrl)
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

    predicted := p.Predict(neurons.FeaturesRow{4.0, 2.0})
    fmt.Printf("predicted 1: %+1.0f, expected: +1\n", predicted)

    predicted = p.Predict(neurons.FeaturesRow{4.0, 1.0})
    fmt.Printf("predicted 2: %+1.0f, expected: -1\n", predicted)
}
