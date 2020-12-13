package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	PREAMBLE_SIZE = 25
)

func main() {
	list := getList()
	n := len(list)
	for i := PREAMBLE_SIZE; i < n; i++ {
		needle := list[i]
		haystack := list[i-PREAMBLE_SIZE : i]
		if !findSum(needle, haystack) {
			fmt.Println(needle)
			contigList := findContiguiusList(needle, list)
			solve(contigList)
		}
	}
}

func getList() []int {
	scanner := getScanner()
	list := []int{}
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		list = append(list, num)
	}
	return list
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

// TODO: Is there a more optimal approach here?
func findSum(needle int, haystack []int) bool {
	n := len(haystack)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if haystack[i]+haystack[j] == needle {
				return true
			}
		}
	}
	return false
}

// TODO: Is there a more optimal approach here?
func findContiguiusList(needle int, haystack []int) []int {
	n := len(haystack)
	for i := 2; i < n; i++ {
		for j := 0; j < n-i; j++ {
			total := 0
			for k := j; k < j+i; k++ {
				total += haystack[k]
			}
			if total == needle {
				return haystack[j : j+i]
			}
		}
	}
	return nil
}

func solve(list []int) {
	min := math.MaxInt32
	max := math.MinInt32

	for _, item := range list {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}
	fmt.Println(min + max)
}
