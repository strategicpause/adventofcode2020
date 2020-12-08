package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := getScanner()
	lines := []string{}
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			total += getAllAnswer(lines)
			lines = []string{}
		} else {
			lines = append(lines, text)
		}
	}
	fmt.Println(total)
}

// Part A Answer
func getAnyAnswer(lines []string) int {
	m := make(map[rune]bool)
	for _, line := range lines {
		for _, char := range line {
			m[char] = true
		}
	}
	return len(m)
}

// Part B Answer
func getAllAnswer(lines []string) int {
	numPeople := len(lines)
	m := make(map[rune]int)
	for _, line := range lines {
		for _, char := range line {
			m[char]++
		}
	}
	allAnswered := 0
	for _, value := range m {
		if value == numPeople {
			allAnswered++
		}
	}
	return allAnswered
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
