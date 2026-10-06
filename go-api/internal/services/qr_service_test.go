package services

import (
	"fmt"
	"testing"
	"math"
)

func TestCalculateQR(t *testing.T) {
	matrix := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	result, err := CalculateQR(matrix)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Q:")
	for _, row := range result.Q {
		fmt.Println(row)
	}

	fmt.Println("R:")
	for _, row := range result.R {
		fmt.Println(row)
	}
}

func TestQRReconstructsMatrix(t *testing.T) {
    matrix := [][]float64{
        {1, 2},
        {3, 4},
        {5, 6},
    }

    result, err := CalculateQR(matrix)

    if err != nil {
        t.Fatal(err)
    }

    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(result.R); j++ {

            value := 0.0

            for k := 0; k < len(result.R); k++ {
                value += result.Q[i][k] * result.R[k][j]
            }

            if math.Abs(value-matrix[i][j]) > 1e-10 {
                t.Fatalf(
                    "expected %v, got %v",
                    matrix[i][j],
                    value,
                )
            }
        }
    }
}