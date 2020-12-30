package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Active    = '.'
	Inactive  = '#'
	NumCycles = 6
	Debug     = true
)

type Vector struct {
	X int
	Y int
	Z int
	W int
}

func (v Vector) getNeighbors() []Vector {
	neighbors := []Vector{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if !(i == 0 && j == 0 && k == 0 && l == 0) {
						vector := Vector{}
						vector.X = v.X + k
						vector.Y = v.Y + j
						vector.Z = v.Z + i
						vector.W = v.W + l
						neighbors = append(neighbors, vector)
					}
				}
			}
		}
	}
	return neighbors
}

func (v Vector) string() string {
	return fmt.Sprintf("X: %d, Y: %d, Z: %d, W: %d", v.X, v.Y, v.Z, v.W)
}

type State map[Vector]bool

func (s State) Cycle() State {
	// Create map which includes all active values and inactive neighbors of those values.
	searchState := State{}
	for vector, active := range s {
		searchState[vector] = active
		for _, neighbor := range vector.getNeighbors() {
			if _, ok := searchState[neighbor]; !ok {
				searchState[neighbor] = false
			}
		}
	}

	newState := State{}
	for vector, active := range searchState {
		activeNeighbors := 0
		for _, neighbor := range vector.getNeighbors() {
			if searchState[neighbor] {
				activeNeighbors++
			}
		}
		if active && (activeNeighbors == 2 || activeNeighbors == 3) {
			newState[vector] = true
		} else if !active && activeNeighbors == 3 {
			newState[vector] = true
		}
	}
	return newState
}

func (s State) GetActiveCubes() int {
	activeCubes := 0
	for _, active := range s {
		if active {
			activeCubes++
		}
	}
	return activeCubes
}

/**
 * Got stuck on this one. Answer influenced by https://github.com/colinodell/advent-2020/blob/main/day17/day17.go
 */
func main() {
	state := getInitialState()
	for i := 0; i < NumCycles; i++ {
		state = state.Cycle()
	}
	activeCubes := state.GetActiveCubes()
	fmt.Println(activeCubes)
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

func getInitialState() State {
	state := State{}
	scanner := getScanner()
	inputState := [][]rune{}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		inputState = append(inputState, line)
	}

	for y, line := range inputState {
		for x, char := range line {
			if char != Active {
				vector := Vector{X: x, Y: y, Z: 0, W: 0}
				state[vector] = true
			}
		}
	}

	return state
}
