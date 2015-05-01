package sudoku

import (
	"bytes"
	"fmt"
	"math/rand"
)

const (
	BoardSize    = 9
	Subdivisions = 3
	PHeight      = BoardSize / Subdivisions
	Height       = BoardSize
	PWidth       = BoardSize / Subdivisions
	Width        = BoardSize

	GuessAlgorithm     = 0
	BackTrackAlgorithm = 1
	DefaultAlgorithm   = BackTrackAlgorithm
)

func Generate() [][]int {
	return GenerateWithAlgorithm(DefaultAlgorithm)
}

func GenerateWithAlgorithm(algorithm int) [][]int {
	return generateFinal(algorithm)
}

func GenerateWithBlanks(numBlanks int) ([][]int, [][]int) {
	return GenerateWithAlgorithmAndBlanks(DefaultAlgorithm, numBlanks)
}

func GenerateWithAlgorithmAndBlanks(algorithm int, numBlanks int) ([][]int, [][]int) {
	final := generateFinal(algorithm)
	retVal := generatePartial(final, numBlanks)
	return retVal, final
}

func Solve(board [][]int) {
	maxX := Width
	maxY := Height
	var firstX = 0
	var firstY = 0
	var needsSolving = false
	for {
		if firstX >= maxX {
			firstX = 0
			firstY++
		}
		if firstY >= maxY {
			break
		}
		if board[firstX][firstY] < 1 {
			needsSolving = true
			break
		}
		firstX++
	}

	if needsSolving {
		solveRecurse(board, firstX, firstY, maxX, maxY)
	}
}

func solveRecurse(board [][]int, x int, y int, maxX int, maxY int) bool {
	var nextX = x + 1
	var nextY = y

	var lastSpace = true
	for {
		if nextX >= maxX {
			nextX = 0
			nextY++
		}
		if nextY >= maxY {
			break
		}
		if board[nextX][nextY] < 1 {
			lastSpace = false
			break
		}
		nextX++
	}

	usableNumbers := goodNumbers(x, y, board)
	length := len(usableNumbers)
	if length < 1 {
		return false
	} else {
		for i := 0; i < length; i++ {
			board[x][y] = usableNumbers[i]
			if lastSpace {
				return true
			} else {
				isSuccess := solveRecurse(board, nextX, nextY, maxX, maxY)
				if isSuccess {
					return true
				}
			}
			board[x][y] = 0
		}
		return false
	}
}

func generateFinal(algorithm int) [][]int {
	switch algorithm {
	case GuessAlgorithm:
		return generateFinal_guessAlgorithm()
	case BackTrackAlgorithm:
		return generateFinal_backtrackAlgorithm()
	default:
		return nil
	}
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
		num = numBlanks
	} else {
		num = numSpaces - numBlanks
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

func generateFinal_guessAlgorithm() [][]int {
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

func generateFinal_backtrackAlgorithm() [][]int {
	retVal := make([][]int, Width)
	for i := 0; i < len(retVal); i++ {
		retVal[i] = make([]int, Height)
	}

	backtrackAlgorithmRecurse(retVal, 0, 0, Width, Height)

	return retVal
}

func backtrackAlgorithmRecurse(retVal [][]int, x int, y int, maxX int, maxY int) bool {
	var usableNumbers = goodNumbers(x, y, retVal)
	length := len(usableNumbers)
	if length < 1 {
		return false
	} else {
		for i := 0; i < length; i++ {
			index := rand.Intn(len(usableNumbers))
			retVal[x][y] = usableNumbers[index]

			var nextX = x + 1
			var nextY = y
			if nextX >= maxX {
				nextX = 0
				nextY++
			}
			if nextY >= maxY {
				return true
			}
			isSuccess := backtrackAlgorithmRecurse(retVal, nextX, nextY, maxX, maxY)
			if isSuccess {
				return true
			} else {
				retVal[x][y] = 0
				usableNumbers = append(usableNumbers[:index], usableNumbers[index+1:]...)
			}
		}
	}
	return false
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

func SudokuString(board [][]int) string {
	buffer := bytes.Buffer{}
	for i := 0; i < len(board); i++ {
		var prefix = ""
		for j := 0; j < len(board[i]); j++ {
			outStr := fmt.Sprintf("%v%d", prefix, board[i][j])
			buffer.WriteString(outStr)
			prefix = " "
		}
		if i != (len(board) - 1) {
			outStr := fmt.Sprintf("\n")
			buffer.WriteString(outStr)
		}
	}
	return buffer.String()
}

func PrintSudoku(board [][]int) {
	printStr := SudokuString(board)
	fmt.Printf(printStr)
}
