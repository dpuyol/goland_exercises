package main

import (
	"fmt"
	"matrix/matrix"
)

func main() {

	matrixResult, err := matrix.New(
		[]float64{2, 7, 8},
		[]float64{4, 4, 7},
		[]float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	matrixResult.Print()

	matrixResult, err = matrix.New(
		[]float64{2, 7, 8},
		[]float64{4, 4, 7, 1},
		[]float64{5, 6, 1})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	matrixResult.Print()

}
