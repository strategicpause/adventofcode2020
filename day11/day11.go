package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EMPTY     = "L"
	OCCUPIED  = "#"
	FLOOR     = "."
	TOLERANCE = 5
	DEBUG     = false
)

func main() {
	state := readInitialState()
	printState(state)
	stateChange := true
	for stateChange {
		state, stateChange = updateState(state)
		printState(state)
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

func printState(state [][]string) {
	if !DEBUG {
		return
	}
	for _, line := range state {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println()
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
			if c == EMPTY && numAdj == 0 {
				newLine = append(newLine, OCCUPIED)
				stateChanged = true
			} else if c == OCCUPIED && numAdj >= TOLERANCE {
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

func getNumAdjacentOccupied(state [][]string, curRow int, curCol int, numRows int, numCols int) int {
	total := 0
	defaultLeft := curCol - 1
	defaultRight := curCol + 1
	defaultTop := curRow - 1
	defaultBottom := curRow + 1
	// Top Left
	for left, top := defaultLeft, defaultTop; left >= 0 && top >= 0; left, top = left-1, top-1 {
		if state[top][left] == OCCUPIED {
			total++
			break
		} else if state[top][left] == EMPTY {
			break
		}
	}
	// Top
	for top := defaultTop; top >= 0; top-- {
		if state[top][curCol] == OCCUPIED {
			total++
			break
		} else if state[top][curCol] == EMPTY {
			break
		}
	}
	// Top Right
	for top, right := defaultTop, defaultRight; top >= 0 && right < numCols; top, right = top-1, right+1 {
		if state[top][right] == OCCUPIED {
			total++
			break
		} else if state[top][right] == EMPTY {
			break
		}
	}
	// Right
	for right := defaultRight; right < numCols; right++ {
		if state[curRow][right] == OCCUPIED {
			total++
			break
		} else if state[curRow][right] == EMPTY {
			break
		}
	}
	// Bottom Right
	for bottom, right := defaultBottom, defaultRight; right < numCols && bottom < numRows; bottom, right = bottom+1, right+1 {
		if state[bottom][right] == OCCUPIED {
			total++
			break
		} else if state[bottom][right] == EMPTY {
			break
		}
	}
	// Bottom
	for bottom := defaultBottom; bottom < numRows; bottom++ {
		if state[bottom][curCol] == OCCUPIED {
			total++
			break
		} else if state[bottom][curCol] == EMPTY {
			break
		}
	}
	// Bottom Left
	for bottom, left := defaultBottom, defaultLeft; left >= 0 && bottom < numRows; bottom, left = bottom+1, left-1 {
		if state[bottom][left] == OCCUPIED {
			total++
			break
		} else if state[bottom][left] == EMPTY {
			break
		}
	}
	// Left
	for left := defaultLeft; left >= 0; left-- {
		if state[curRow][left] == OCCUPIED {
			total++
			break
		} else if state[curRow][left] == EMPTY {
			break
		}
	}
	return total
}
