package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AngelSalazarV/talsory-challenge/go-api/internal/models"
)

type NodeStatsResponse struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Average    float64 `json:"average"`
	Total      float64 `json:"total"`
	IsDiagonal bool    `json:"isDiagonal"`
}

func SendMatrixToNode(matrix [][]float64) (NodeStatsResponse, error) {
	requestBody := models.MatrixRequest{
		Matrix: matrix,
	}

	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		return NodeStatsResponse{}, err
	}

	nodeAPIURL := os.Getenv("NODE_API_URL")

	if nodeAPIURL == "" {
			nodeAPIURL = "http://localhost:3001"
	}

	response, err := http.Post(
			nodeAPIURL+"/api/stats",
			"application/json",
			bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return NodeStatsResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return NodeStatsResponse{}, fmt.Errorf(
			"node api returned status %d",
			response.StatusCode,
		)
	}

	var result NodeStatsResponse

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return NodeStatsResponse{}, err
	}

	return result, nil
}