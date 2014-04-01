package sudoku

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRegular(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output := Generate()
	fmt.Printf("%v\n", output)
}

func TestWithGaps(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	output, answer := GenerateWithBlanks(30)
	fmt.Printf("Answer: %v\nOutput: %v", answer, output)
}
