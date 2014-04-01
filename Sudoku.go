// Sudoku project Sudoku.go
package Sudoku

import (
	"math/rand"
)

const (
	boardSize        = 9
	subdivisions     = 3
	pHeight          = boardSize / subdivisions
	height           = boardSize
	pWidth           = boardSize / subdivisions
	width            = boardSize
	DefaultNumBlanks = 30
)

func Generate() [][]int {
	return GenerateWithBlanks(DefaultNumBlanks)
}

func GenerateWithBlanks(numBlanks int) [][]int {
	final := generateFinal()
	retVal := generatePartial(final, numBlanks)
	return retVal
}

func generatePartial(final [][]int, numBlanks int) [][]int {

	var numSpaces = 0
	retVal := make([][]int, len(final))
	for i := 0; i < len(final); i++ {
		retVal[i] = make([]int, len(final[i]))
		numSpaces += len(final[i])
	}

	fillWithNulls := numBlanks < (numSpaces / 2)

	if fillWithNulls {
		for i := 0; i < len(retVal); i++ {
			copy(retVal[i], final[i])
		}
	}

	for i := 0; i < numBlanks; {
		randX := rand.Intn(len(retVal))
		randCol := retVal[randX]
		randY := rand.Intn(len(randCol))
		randVal := randCol[randY]
		if fillWithNulls {
			if randVal > 0 {
				retVal[randX][randY] = 0
				i++
			}
		} else {
			if randVal < 1 {
				retVal[randX][randY] = final[randX][randY]
				i++
			}
		}
	}
	return retVal
}

func generateFinal() [][]int {
	retVal := make([][]int, width)
	for i := 0; i < width; i++ {
		retVal[i] = make([]int, height)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			usableNumbers := goodNumbers(i, j, retVal)
			retVal[i][j] = usableNumbers[rand.Intn(len(usableNumbers))]
		}
	}

	return retVal
}

func goodNumbers(x int, y int, currentBoard [][]int) []int {
	// TODO handle holders
	var squareHolder = make([]int, 0, boardSize)
	for i := 1; i <= boardSize; i++ {
		squareHolder = append(squareHolder, i)
	}
	var lineHolderV = make([]int, 0, boardSize)
	for i := 1; i <= boardSize; i++ {
		lineHolderV = append(lineHolderV, i)
	}
	lineHolderH := make([]int, 0, boardSize)
	for i := 1; i <= boardSize; i++ {
		lineHolderH = append(lineHolderH, i)
	}

	var retVal = make([]int, 0, boardSize)
	for i := 1; i <= boardSize; i++ {
		var useValue = true
		for j := 0; j < len(squareHolder); j++ {
			if squareHolder[j] == i {
				useValue = false
				break
			}
		}
		if useValue {
			for j := 0; j < len(lineHolderV); j++ {
				if lineHolderV[j] == i {
					useValue = false
					break
				}
			}
		}
		if useValue {
			for j := 0; j < len(lineHolderH); j++ {
				if lineHolderH[j] == i {
					useValue = false
					break
				}
			}
		}

		if useValue {
			retVal = append(retVal, i)
		}
	}
	return retVal
}
