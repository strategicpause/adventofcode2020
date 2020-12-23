package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	EAST         = 0
	SOUTH        = 90
	WEST         = 180
	NORTH        = 270
	MOVE_NORTH   = "N"
	MOVE_SOUTH   = "S"
	MOVE_EAST    = "E"
	MOVE_WEST    = "W"
	TURN_LEFT    = "L"
	TURN_RIGHT   = "R"
	MOVE_FORWARD = "F"
)

func main() {
	x, y := 10, 1
	shipX, shipY := 0, 0
	scanner := getScanner()
	for scanner.Scan() {
		text := scanner.Text()
		action := text[:1]
		units, _ := strconv.Atoi(text[1:])
		switch action {
		case TURN_LEFT:
			radians := float64(units) * (math.Pi / 180.0)
			newX := int(float64(x)*math.Cos(radians)) - int(float64(y)*math.Sin(radians))
			newY := int(float64(x)*math.Sin(radians)) + int(float64(y)*math.Cos(radians))
			x, y = newX, newY
			break
		case TURN_RIGHT:
			radians := float64(360-units) * (math.Pi / 180.0)
			newX := int(float64(x)*math.Cos(radians)) - int(float64(y)*math.Sin(radians))
			newY := int(float64(x)*math.Sin(radians)) + int(float64(y)*math.Cos(radians))
			x, y = newX, newY
			break
		case MOVE_FORWARD:
			for i := 0; i < units; i++ {
				shipX += x
				shipY += y
			}
			break
		default:
			x, y = move(action, units, x, y)
		}
	}
	distance := math.Abs(float64(shipX)) + math.Abs(float64(shipY))
	fmt.Println(distance)
}

func move(action string, units int, x int, y int) (int, int) {
	switch action {
	case MOVE_NORTH:
		return x, y + units
	case MOVE_SOUTH:
		return x, y - units
	case MOVE_EAST:
		return x + units, y
	case MOVE_WEST:
		return x - units, y
	}
	return x, y
}

func getActionFromDirection(direction int) string {
	switch direction {
	case NORTH:
		return MOVE_NORTH
	case SOUTH:
		return MOVE_SOUTH
	case EAST:
		return MOVE_EAST
	case WEST:
		return MOVE_WEST
	}
	fmt.Println(direction)
	return ""
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
