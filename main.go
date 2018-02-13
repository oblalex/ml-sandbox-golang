package main

import "fmt"
import "elements"


func main() {
    p := elements.NewPerceptron(2, 1000, 0.01)
    fmt.Printf(
        "features number = %d, iterations number = %d, learning step = %f\n",
        p.FeaturesNumber,
        p.IterationsNumber,
        p.LearningStep,
    )
}
