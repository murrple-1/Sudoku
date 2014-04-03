package sudoku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	LowBlanks  = 1
	MidBlanks  = 30
	HighBlanks = 80
)

func printOutput(output [][]int) {
	fmt.Printf("Output:\n")
	PrintSudoku(output)
	fmt.Printf("\n")
}

func printAnswerAndOutput(answer [][]int, output [][]int) {
	fmt.Printf("Answer:\n")
	PrintSudoku(answer)
	fmt.Printf("\n")
	printOutput(output)
}

func countSudokuSquares(output [][]int) (int, int) {
	var blanks = 0
	var filled = 0
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			if output[i][j] > 0 {
				filled++
			} else {
				blanks++
			}
		}
	}
	return blanks, filled
}

func TestStandard(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output := Generate()
	printOutput(output)
}

func TestStandardWithLowBlanks(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithBlanks(LowBlanks)
	printAnswerAndOutput(answer, output)
	blanks, _ := countSudokuSquares(output)

	if blanks != LowBlanks {
		t.Errorf("Incorrect Numer of Blanks: %d - Expected: %d", blanks, LowBlanks)
	}
}

func TestStandardWithMidBlanks(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithBlanks(MidBlanks)
	printAnswerAndOutput(answer, output)
	blanks, _ := countSudokuSquares(output)

	if blanks != MidBlanks {
		t.Errorf("Incorrect Numer of Blanks: %d - Expected: %d", blanks, MidBlanks)
	}
}

func TestStandardWithHighBlanks(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithBlanks(HighBlanks)
	printAnswerAndOutput(answer, output)
	blanks, _ := countSudokuSquares(output)

	if blanks != HighBlanks {
		t.Errorf("Incorrect Numer of Blanks: %d - Expected: %d", blanks, HighBlanks)
	}
}

func TestGenerate_guessAlgorithm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output := GenerateWithAlgorithm(GuessAlgorithm)
	printOutput(output)
}

func TestGenerateWithBlanks_guessAlgorithm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithAlgorithmAndBlanks(GuessAlgorithm, MidBlanks)
	printAnswerAndOutput(answer, output)
	blanks, _ := countSudokuSquares(output)

	if blanks != MidBlanks {
		t.Errorf("Incorrect Numer of Blanks: %d - Expected: %d", blanks, MidBlanks)
	}
}

func TestGenerate_backtrackAlgorithm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output := GenerateWithAlgorithm(BackTrackAlgorithm)
	printOutput(output)
}

func TestGenerateWithBlanks_backtrackAlgorithm(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithAlgorithmAndBlanks(BackTrackAlgorithm, MidBlanks)
	printAnswerAndOutput(answer, output)
	blanks, _ := countSudokuSquares(output)

	if blanks != MidBlanks {
		t.Errorf("Incorrect Numer of Blanks: %d - Expected: %d", blanks, MidBlanks)
	}
}
