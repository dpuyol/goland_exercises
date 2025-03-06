package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct {
	matrix    [][]float64
	height    int64
	weight    int64
	quadratic bool
}

func New(vector ...[]float64) (*Matrix, error) {
	matrix := Matrix{
		height: int64(len(vector)),
	}

	for i := range vector {

		if i == 0 {
			matrix.weight = int64(len(vector[i]))
		} else if matrix.weight != int64(len(vector[i])) {
			return nil, fmt.Errorf("Error in matrix size: expected %d, got %d", matrix.weight, len(vector[i]))
		}
	}

	matrix.matrix = vector
	matrix.quadratic = matrix.height == matrix.weight
	return &matrix, nil
}

func (matrix *Matrix) Print() {

	var sb strings.Builder
	for i := range matrix.matrix {

		sb.WriteString("[ ")
		for j, val := range matrix.matrix[i] {

			if j != 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(fmt.Sprintf("%v", val))
		}
		sb.WriteString(" ]\n")
	}

	sb.WriteString(fmt.Sprintf("%dx%d\n", matrix.height, matrix.weight))
	sb.WriteString(fmt.Sprintf("Is quadratic: %v\n", matrix.quadratic))

	fmt.Print(sb.String())
}
