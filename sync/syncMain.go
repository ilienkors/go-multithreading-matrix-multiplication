package main

import (
	"fmt"
	"time"
)

const n = 2560

func rowMultiplication(rowNumber int, matrixA, matrixB, matrixRes *[n][n]int) {
	var sum int
	for i := 0; i < len(matrixA); i++ {
		sum = 0
		for j := 0; j < len(matrixA); j++ {
			sum += matrixA[rowNumber][j] * matrixB[j][i]
		}
		matrixRes[rowNumber][i] = sum
	}
}

func matrixGen(matrixA, matrixB, matrixRes *[n][n]int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrixA[i][j] = 3
			matrixB[i][j] = 3
			matrixRes[i][j] = 0
		}
	}
}

func main() {
	var matrixA, matrixB, matrixRes [n][n]int
	matrixGen(&matrixA, &matrixB, &matrixRes)

	start := time.Now()
	for i := 0; i < len(matrixA); i++ {
		rowMultiplication(i, &matrixA, &matrixB, &matrixRes)
	}
	fmt.Print(time.Since(start), "\n")
}
