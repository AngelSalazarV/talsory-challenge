package services

import (
    "errors"

    "github.com/AngelSalazarV/talsory-challenge/go-api/internal/models"
)

func CalculateQR(matrix [][]float64) (models.QRResponse, error) {
    if len(matrix) == 0 {
        return models.QRResponse{}, errors.New("matrix cannot be empty")
    }

    return models.QRResponse{}, nil
}