package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Square rune = '.'
	Tree   rune = '#'
	Rows        = 323
	Cols        = 31
)

func main() {
	partA()
	partB()
}

func partA() {
	geoMap := buildGeoMap()
	trees := getTreesForTraverseStrategy(geoMap, 1, 3)
	fmt.Println(trees)
}

func partB() {
	geoMap := buildGeoMap()
	strategies := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	trees := 1
	for _, strategy := range strategies {
		trees *= getTreesForTraverseStrategy(geoMap, strategy[0], strategy[1])
	}
	fmt.Println(trees)
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

func buildGeoMap() *[Rows][Cols]rune {
	scanner := getScanner()
	geoMap := [Rows][Cols]rune{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, char := range line {
			geoMap[row][col] = char
		}
		row++
	}
	return &geoMap
}

func getTreesForTraverseStrategy(geoMap *[Rows][Cols]rune, rowsToTravel int, colsToTravel int) int {
	curCol, curRow := 0, 0
	trees := 0
	for curRow < Rows-1 {
		curRow += rowsToTravel
		curCol += colsToTravel
		if (*geoMap)[curRow%Rows][curCol%Cols] == Tree {
			trees++
		}
	}
	return trees
}
