package main

import (
	"fmt"
	"math"
)

func multiply(a, b float64) float64 {
	return a * b
}

func quadRoot(a, b, c float64) (result []float64) {

	if a == 0 {
		result = append(result, c/b)
		return
	}

	posRoot := b * -1
	posRoot = posRoot + math.Sqrt(math.Pow(b, 2)-(4*(a*c)))
	posRoot = posRoot / multiply(2, a)

	negRoot := b * -1
	negRoot = negRoot - math.Sqrt(math.Pow(b, 2)-(4*(a*c)))
	negRoot = negRoot / multiply(2, a)

	result = append(result, negRoot)
	result = append(result, posRoot)

	return
}

func main() {
	fmt.Println(quadRoot(7, 8, -11))
	fmt.Println(quadRoot(0, 8, -11))
}
