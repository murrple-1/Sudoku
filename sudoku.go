package sudoku

import (
	"math/rand"
)

const (
	BoardSize    = 9
	Subdivisions = 3
	PHeight      = BoardSize / Subdivisions
	Height       = BoardSize
	PWidth       = BoardSize / Subdivisions
	Width        = BoardSize
)

func Generate() [][]int {
	return generateFinal()
}

func GenerateWithBlanks(numBlanks int) ([][]int, [][]int) {
	final := generateFinal()
	retVal := generatePartial(final, numBlanks)
	return retVal, final
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

	var num int
	if fillWithNulls {
		num = numSpaces - numBlanks
	} else {
		num = numBlanks
	}

	for i := 0; i < num; {
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
	var success = false
	var retVal [][]int
	for !success {
		success = true
		retVal = make([][]int, Width)
		for i := 0; i < len(retVal); i++ {
			retVal[i] = make([]int, Height)
		}

		for i := 0; i < len(retVal); i++ {
			for j := 0; j < len(retVal[i]); j++ {
				usableNumbers := goodNumbers(i, j, retVal)
				if len(usableNumbers) < 1 {
					success = false
					goto error
				}
				retVal[i][j] = usableNumbers[rand.Intn(len(usableNumbers))]
			}
		}
	error:
	}

	return retVal
}

func goodNumbers(x int, y int, currentBoard [][]int) []int {
	var bad = make([]int, 0, BoardSize*3)

	{
		var minX = 0
		var minY = 0
		for i := 0; i < BoardSize; i += (BoardSize / Subdivisions) {
			if x >= i && x < (i+(BoardSize/Subdivisions)) {
				minX = i
				break
			}
		}
		for i := 0; i < BoardSize; i += (BoardSize / Subdivisions) {
			if y >= i && y < (i+(BoardSize/Subdivisions)) {
				minY = i
				break
			}
		}
		for i := minX; i < (minX + (BoardSize / Subdivisions)); i++ {
			for j := minY; j < (minY + (BoardSize / Subdivisions)); j++ {
				val := currentBoard[i][j]
				if val > 0 {
					bad = append(bad, val)
				}
			}
		}
	}

	for i := 0; i < BoardSize; i++ {
		val := currentBoard[i][y]
		if val > 0 {
			bad = append(bad, val)
		}
	}

	for i := 0; i < BoardSize; i++ {
		val := currentBoard[x][i]
		if val > 0 {
			bad = append(bad, val)
		}
	}

	var retVal = make([]int, 0, BoardSize)
	for i := 1; i <= BoardSize; i++ {
		retVal = append(retVal, i)
	}

	for i := 0; i < len(bad); i++ {
		badVal := bad[i]
		var j = 0
		var isBad = false
		for ; j < len(retVal); j++ {
			if badVal == retVal[j] {
				isBad = true
				break
			}
		}

		if isBad {
			retVal = append(retVal[:j], retVal[j+1:]...)
		}
	}
	return retVal
}
