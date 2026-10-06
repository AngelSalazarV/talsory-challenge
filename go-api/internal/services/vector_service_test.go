package services

import "testing"

func TestDotProduct(t *testing.T) {
    a := []float64{1, 2, 3}
    b := []float64{4, 5, 6}

    result := dotProduct(a, b)

    expected := 32.0

    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}

func TestScalarMultiply(t *testing.T) {
    scalar := 2.0
    vector := []float64{1, 2, 3}

    result := scalarMultiply(scalar, vector)

    expected := []float64{2, 4, 6}

    for i := 0; i < len(expected); i++ {
        if result[i] != expected[i] {
            t.Errorf("expected %v, got %v", expected, result)
        }
    }
}

func TestSubtractVectors(t *testing.T) {
    a := []float64{5, 7, 9}
    b := []float64{1, 2, 3}

    result := subtractVectors(a, b)

    expected := []float64{4, 5, 6}

    for i := 0; i < len(expected); i++ {
        if result[i] != expected[i] {
            t.Errorf("expected %v, got %v", expected, result)
        }
    }
}

func TestVectorNorm(t *testing.T) {
    vector := []float64{3, 4}

    result := vectorNorm(vector)

    expected := 5.0

    if result != expected {
        t.Errorf("expected %v, got %v", expected, result)
    }
}

func TestNormalizeVector(t *testing.T) {
    vector := []float64{3, 4}

    result := normalizeVector(vector)

    expected := []float64{0.6, 0.8}

    for i := 0; i < len(expected); i++ {
        if result[i] != expected[i] {
            t.Errorf("expected %v, got %v", expected, result)
        }
    }
}