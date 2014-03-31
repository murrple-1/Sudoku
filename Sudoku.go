// Sudoku project Sudoku.go
package Sudoku

import ()

const (
	pHeight = 3
	height  = pHeight * 3
	pWidth  = 3
	width   = pWidth * 3
)

func Generate() [][]int8 {
	final := generateFinal()
	retVal := generatePartial(final)
	return retVal
}

func generatePartial(final [][]int8) [][]int8 {
	retVal := make([][]int8, width)
	for i := 0; i < width; i++ {
		retVal[i] = make([]int8, height)
		for j := 0; j < height; j++ {
			retVal[i][j] = -1
		}
	}
	return retVal
}

func generateFinal() [][]int8 {
	retVal := make([][]int8, width)
	for i := 0; i < width; i++ {
		retVal[i] = make([]int8, height)
		for j := 0; j < width; j++ {
			retVal[i][j] = -1
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// TODO usableNumbers :=
		}
	}

	return retVal
}
