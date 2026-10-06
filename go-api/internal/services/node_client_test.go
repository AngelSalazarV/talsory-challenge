package services

import (
	"testing"
)

func TestSendMatrixToNode(t *testing.T) {
	matrix := [][]float64{
		{1, 2},
		{3, 4},
	}

	result, err := SendMatrixToNode(matrix)

	if err != nil {
		t.Fatal(err)
	}

	if result.Max != 4 {
		t.Errorf("expected max 4, got %v", result.Max)
	}

	if result.Min != 1 {
		t.Errorf("expected min 1, got %v", result.Min)
	}

	if result.Average != 2.5 {
		t.Errorf("expected average 2.5, got %v", result.Average)
	}

	if result.Total != 10 {
		t.Errorf("expected total 10, got %v", result.Total)
	}

	if result.IsDiagonal {
		t.Errorf("expected matrix to not be diagonal")
	}
}