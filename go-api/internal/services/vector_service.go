package services

import "math"

func dotProduct(a []float64, b []float64)float64 {
	result := 0.0

	for i:=0; i<len(a); i++{
		result += a[i]*b[i]
	}

	return result
}

func scalarMultiply(scalar float64, vector []float64) []float64 {
	result := make([]float64, len(vector))

	for i:=0; i<len(vector); i++ {
		result[i] = scalar*vector[i]
	} 

	return result
}

func subtractVectors(a []float64, b []float64)[] float64 {
	result := make([]float64, len(a))

	for i:=0; i<len(a); i++ {
		result[i] = a[i]-b[i]
	}

	return result
}

func vectorNorm(vector []float64)float64 {
	sum := 0.0

	for i:=0; i<len(vector); i++ {
		sum += vector[i] * vector[i]
	}

	return math.Sqrt(sum)
}

func normalizeVector(vector []float64) []float64 {
	norm := vectorNorm(vector)

	result := make([]float64, len(vector))

	for i:=0; i<len(vector); i++ {
		result[i] = vector[i] / norm
	}	

	return result
}

func TestExample() {
    q1 := []float64{0.169, 0.507, 0.845}
    a2 := []float64{2, 4, 6}

    r12 := dotProduct(q1, a2)

    projection := scalarMultiply(r12, q1)

    v2 := subtractVectors(a2, projection)

    q2 := normalizeVector(v2)

    println("r12:", r12)
    println("q2:", q2[0], q2[1], q2[2])
}