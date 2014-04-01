Sudoku
======

Experimental Sudoku generator.

The algorithm used is very simple and naive. There are faster and more concise algorithms, but this one will probably generate a complete Sudoku board relatively quickly.

It works by moving through each slot and randomly chosing an available number for that slot. When there are no numbers available, the generation has failed and it tries again from the top. Eventually, it will generate a complete Sudoku board. Due to the nondeterministic approach this algorithm takes, it is not possible to determine how long it will take (theoretically, it could never return), but for the benchmark runs, it returns in under a second on modern machines.

To generate a complete puzzle, use:
```golang
sudoku.Generate() [][]int
```

To return a puzzle with slots missing, use:
```golang
sudoku.GenerateWithBlanks(numBlanks int) ([][]int, [][]int) //return 1: Blanked Puzzle, return 2: Complete puzzle
```
