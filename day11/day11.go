package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EMPTY    = "L"
	OCCUPIED = "#"
	FLOOR    = "."
)

func main() {
	state := readInitialState()
	stateChange := true
	for stateChange {
		state, stateChange = updateState(state)
	}
	numOccupied := getNumOccupied(state)
	fmt.Println(numOccupied)
}

func getNumOccupied(state [][]string) int {
	numOccupied := 0
	for _, line := range state {
		for _, c := range line {
			if c == OCCUPIED {
				numOccupied++
			}
		}
	}
	return numOccupied
}

func readInitialState() [][]string {
	state := [][]string{}
	scanner := getScanner()
	for scanner.Scan() {
		s := scanner.Text()
		state = append(state, strings.Split(s, ""))
	}
	return state
}

func getScanner() *bufio.Scanner {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func updateState(previousState [][]string) ([][]string, bool) {
	numRows := len(previousState)
	numCols := len(previousState[0])
	newState := [][]string{}
	stateChanged := false
	for curRow, line := range previousState {
		newLine := []string{}
		for curCol, c := range line {
			numAdj := getNumAdjacentOccupied(previousState, curRow, curCol, numRows, numCols)
			if c == FLOOR {
				newLine = append(newLine, FLOOR)
			} else if c == EMPTY && numAdj == 0 {
				newLine = append(newLine, OCCUPIED)
				stateChanged = true
			} else if c == OCCUPIED && numAdj > 3 {
				newLine = append(newLine, EMPTY)
				stateChanged = true
			} else {
				newLine = append(newLine, c)
			}
		}
		newState = append(newState, newLine)
	}
	return newState, stateChanged
}

/**
 * i - current row
 * j - current col
 * m - row length
 * n - col length
 */
func getNumAdjacentOccupied(state [][]string, curRow int, curCol int, numRows int, numCols int) int {
	total := 0
	left := curCol - 1
	right := curCol + 1
	top := curRow - 1
	bottom := curRow + 1
	// Top Left
	if left >= 0 && top >= 0 && state[top][left] == OCCUPIED {
		total++
	}
	// Top
	if top >= 0 && state[top][curCol] == OCCUPIED {
		total++
	}
	// Top Right
	if right < numCols && top >= 0 && state[top][right] == OCCUPIED {
		total++
	}
	// Right
	if right < numCols && state[curRow][right] == OCCUPIED {
		total++
	}
	// Bottom Right
	if right < numCols && bottom < numRows && state[bottom][right] == OCCUPIED {
		total++
	}
	// Bottom
	if bottom < numRows && state[bottom][curCol] == OCCUPIED {
		total++
	}
	// Bottom Left
	if left >= 0 && bottom < numRows && state[bottom][left] == OCCUPIED {
		total++
	}
	// Left
	if left >= 0 && state[curRow][left] == OCCUPIED {
		total++
	}
	return total
}
