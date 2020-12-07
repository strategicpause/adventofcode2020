package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	partA()
	partB()
}

func partA() {
	scanner := getScanner()
	highestSeatId := -1
	for scanner.Scan() {
		seatCode := scanner.Text()
		seatId := getSeatId(seatCode)
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}
	fmt.Println(highestSeatId)
}

func partB() {
	seats := [884]bool{}
	scanner := getScanner()
	for scanner.Scan() {
		seatCode := scanner.Text()
		seatId := getSeatId(seatCode)
		seats[seatId] = true
	}
	n := len(seats)
	for i := 0; i < n; i++ {
		if !seats[i] {
			fmt.Println(i)
		}
	}

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

func getSeatId(seatCode string) int {
	minRow, maxRow := 0, 127
	for i := 0; i < 7; i++ {
		if seatCode[i] == 'F' {
			maxRow = (minRow + maxRow) / 2
		} else if seatCode[i] == 'B' {
			minRow = (minRow + maxRow) / 2
		}
	}
	row := maxRow
	col := 0.0
	for i := 0; i <= 2; i++ {
		if seatCode[7+i] == 'R' {
			col += math.Pow(2, 2.0-float64(i))
		}
	}
	return row*8 + int(col)
}
