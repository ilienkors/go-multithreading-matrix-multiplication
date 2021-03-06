package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 2560

func rowMultiplication(rowNumber int, matrixA, matrixB, matrixRes *[n][n]int, wg *sync.WaitGroup) {
	defer wg.Done()
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
	wg := new(sync.WaitGroup)
	wg.Add(n)

	var matrixA, matrixB, matrixRes [n][n]int
	matrixGen(&matrixA, &matrixB, &matrixRes)

	start := time.Now()
	for i := 0; i < len(matrixA); i++ {
		go rowMultiplication(i, &matrixA, &matrixB, &matrixRes, wg)
	}
	wg.Wait()

	fmt.Print(time.Since(start), "\n")
}
