package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrixReturn := [][]int{}
	for x := 0; x < rows; x++ {
		if len(matrixReturn) <= x {
			matrixReturn = append(matrixReturn, make([][]int, 1)...)
		}
		for y := 0; y < cols; y++ {
			if len(matrixReturn[x]) <= y {
				matrixReturn[x] = append(matrixReturn[x], make([]int, 1)...)
			}
			matrixReturn[x][y] = x * y
		}
	}
	return matrixReturn
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
