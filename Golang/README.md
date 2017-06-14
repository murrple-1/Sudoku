Sudoku
======
This library generates Sudoku puzzles in golang. By default, it uses the standard backtrack algorithm to quickly create valid boards. However, there is also an option to use an experimental non-deterministic algorithm that I developed.

## Backtrack Algorithm
It works through a recursive algorithm, which will ensure that if a dead-end is hit in the generation, it will roll back up the recursion, trying new numbers and continuing to back track until it succeeds.

## Guess Algorithm
It works by moving through each slot and randomly chosing an available number for that slot. When there are no numbers available, the generation has failed and it tries again from the top. Eventually, it will generate a complete Sudoku board. Due to the nondeterministic approach this algorithm takes, it is not possible to determine how long it will take (theoretically, it could never return), but for the benchmark runs, it returns in under a second on modern machines.

### Code Examples
#### Generation
To generate a complete puzzle, use:
```go
func sudoku.Generate() [][]int
```

To return a puzzle with slots missing, use:
```go
func sudoku.GenerateWithBlanks(numBlanks int) ([][]int, [][]int) //return 1: Blanked Puzzle, return 2: Complete puzzle
```
To determine the algorithm to use, use:
```go
func sudoku.GenerateWithAlgorithm(algorithm int) [][]int
func sudoku.GenerateWithAlgorithmAndBlanks(algorithm int, numBlanks int) ([][]int, [][]int)
```
The algorithm constants are:
```go
const (
  GuessAlgorithm = 0
  BacktrackAlgorithm = 1
  )
```
#### Solver
To solve a puzzle, use:
```go
func sudoku.Solve(board [][]int)
```
#### Helpers
Also included are some helpful methods to print the Sudoku to console or log:
```go
func sudoku.SudokuString(board [][]int) string
func sudoku.PrintSudoku(board [][]int)
```
