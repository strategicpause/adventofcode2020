package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	partA()
	partB()
}

func partA() {
	scanner := getScanner()

	valid := 0

	for scanner.Scan() {
		text := scanner.Text()
		password, letter, min, max := parse(text)

		count := 0
		for _, c := range password {
			if string(c) == letter {
				count++
			}
		}
		if !(count < min || count > max) {
			valid++
		}
	}
	fmt.Println(valid)
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

func parse(text string) (password string, letter string, min int, max int) {
	s := strings.Split(text, ": ")
	rule, password := s[0], s[1]
	s = strings.Split(rule, " ")
	constraints, letter := s[0], s[1]
	s = strings.Split(constraints, "-")
	min, _ = strconv.Atoi(s[0])
	max, _ = strconv.Atoi(s[1])
	return password, letter, min, max
}

func partB() {
	scanner := getScanner()
	valid := 0
	for scanner.Scan() {
		text := scanner.Text()
		password, letter, min, max := parse(text)

		n := len(password)
		hasFirst := min <= n && string(password[min-1]) == letter
		hasSecond := max <= n && string(password[max-1]) == letter
		if hasFirst != hasSecond {
			valid++
		}
	}
	fmt.Println(valid)
}
