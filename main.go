package main

import (
	"fmt"
)

func rowMultiplication(rowNumber int, row []int, matrixB [][]int, c chan [][]int) {
	sumRow := [][]int{
		[]int{rowNumber},
		[]int{},
	}

	for i := 0; i < len(row); i++ {
		sum := 0
		for j := 0; j < len(row); j++ {
			sum += row[j] * matrixB[j][i]
		}
		sumRow[1] = append(sumRow[1], sum)
	}
	c <- sumRow
}

func main() {
	c := make(chan [][]int)

	matrixA := [][]int{
		[]int{9, 2, 3},
		[]int{9, 6, 3},
		[]int{9, 7, 3},
	}

	matrixB := [][]int{
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
	}

	matrixC := [][]int{
		[]int{},
		[]int{},
		[]int{},
	}


	for i := 0; i < len(matrixA); i++ {
		go rowMultiplication(i, matrixA[i], matrixB, c)
	}

	for i := 0; i < len(matrixA); i++ {
		rowInfo := <-c 
		matrixC[rowInfo[0][0]] = rowInfo[1]
	}
	fmt.Print(matrixC)

}
