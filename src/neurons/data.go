package neurons

import (
    "encoding/csv"
    "path/filepath"
    "os"
    "strconv"
)

type FeaturesRow []float64


type LabeledFeaturesRow struct {
    Features FeaturesRow
    Label    float64
}


type LabeledFeaturesSeries []LabeledFeaturesRow


func StoreLabeledFeaturesSeries(
    series   LabeledFeaturesSeries,
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

    for _, row := range series {
        data := []string{
            strconv.FormatFloat(row.Features[0], 'f', -1, 64),
            strconv.FormatFloat(row.Features[1], 'f', -1, 64),
            strconv.FormatFloat(row.Label,       'f', -1, 64),
        }

        err := writer.Write(data)
        if err != nil {
            return err
        }
    }

    return nil
}
