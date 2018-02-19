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


const (
    DATA_SET_URL = "https://archive.ics.uci.edu/ml/machine-learning-databases/iris/iris.data"
    DATA_DIR_PATH = "./data"
    OUTPUT_FILE_PATH = DATA_DIR_PATH + "/" + "training-set"
    CACHE_DIR_PATH = "./.cache"
)


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
    fileName := getURLHash(url)
    filePath := fmt.Sprintf("%s/%s", CACHE_DIR_PATH, fileName)

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


func buildTrainingSet(data [][]string) ([]neurons.TrainingSetRow, error) {
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


func storeTrainingSet(
    trainingSet []neurons.TrainingSetRow,
    filePath string,
) (error) {
    dirPath := filepath.Dir(filePath)
    if _, err := os.Stat(dirPath); os.IsNotExist(err) {

        err := os.Mkdir(dirPath, 0755)
        if err != nil {
            return err
        }
    }

    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, row := range trainingSet {
        data := []string{
            strconv.FormatFloat(row.Features[0], 'f', -1, 64),
            strconv.FormatFloat(row.Features[1], 'f', -1, 64),
            strconv.FormatFloat(row.Expected,    'f', -1, 64),
        }

        err := writer.Write(data)
        if err != nil {
            return err
        }
    }

    return nil
}


func main() {
    raw, err := maybeDownload(DATA_SET_URL)
    if err != nil {
        panic(err)
    }

    strings, err := readCSVdata(bytes.NewReader(raw))
    if err != nil {
        panic(err)
    }

    trainingSet, err := buildTrainingSet(strings)
    if err != nil {
        panic(err)
    }

    err = storeTrainingSet(trainingSet, OUTPUT_FILE_PATH)
    if err != nil {
        panic(err)
    }
}
