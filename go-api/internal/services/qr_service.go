package services

import (
    "errors"

    "github.com/AngelSalazarV/talsory-challenge/go-api/internal/models"
)

func getColumn(matrix [][]float64, column int) []float64 {
	result := make([]float64, len(matrix))

	for i:=0; i<len(matrix); i++ {
		result[i] = matrix[i][column]
	}
	
	return result
}

func CalculateQR(matrix [][]float64) (models.QRResponse, error) {
    if len(matrix) == 0 {
        return models.QRResponse{}, errors.New("matrix cannot be empty")
    }

		columns := len(matrix[0])

		for i:=1; i<len(matrix);i++ {
			if len(matrix[i]) != columns {
				return models.QRResponse{}, errors.New("matrix must be rectangular")
			}
		}

		if len(matrix) < columns {
			return models.QRResponse{}, errors.New("matrix must have rows >= columns")
		}

    rows := len(matrix)

		q := make([][]float64, rows)
		for i := 0; i < rows; i++ {
				q[i] = make([]float64, columns)
		}

		r := make([][]float64, columns)
		for i := 0; i < columns; i++ {
				r[i] = make([]float64, columns)
		}

		for j:=0; j<columns; j++ {
			column := getColumn(matrix, j)

			for i:=0; i<j; i++ {	
				qcolumn := getColumn(q, i)
				r[i][j] = dotProduct(qcolumn, column)
				projection := scalarMultiply(r[i][j], qcolumn)
				column = subtractVectors(column, projection)
			}
			norm := vectorNorm(column)
			r[j][j] = norm
			qcolumn := normalizeVector(column)

			for i:=0; i<rows; i++ {
				q[i][j] = qcolumn[i]
			}
		}

		return models.QRResponse{
				Q: q,
				R: r,
		}, nil
}