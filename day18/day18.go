package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	arr []string
}

func (s *Stack) Push(e string) {
	s.arr = append(s.arr, e)
}

func (s *Stack) Pop() string {
	n := len(s.arr) - 1
	e := s.arr[n]
	s.arr = s.arr[:n]
	return e
}

func (s *Stack) Peek() string {
	n := len(s.arr) - 1
	if n < 0 {
		return ""
	}
	return s.arr[n]
}

func (s *Stack) Empty() bool {
	return len(s.arr) == 0
}

func main() {
	scanner := getScanner()
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		value := evaluateExpression(line)
		fmt.Println(value)
		total += value
	}
	fmt.Println(total)
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

func evaluateExpression(line string) int {
	stack := Stack{}
	input := strings.Split(line, " ")
	for _, c := range input {
		if c == "*" || c == "+" || c == "(" {
			stack.Push(c)
		} else if c == ")" {
			temp := stack.Pop()
			stack.Pop() // Remove "("
			if stack.Peek() == "*" || stack.Peek() == "+" {
				stack.Push(temp)
				answer := evaluate(&stack)
				stack.Push(answer)
			} else {
				stack.Push(temp)
			}
		} else if stack.Empty() || stack.Peek() == "(" {
			stack.Push(c)
		} else {
			stack.Push(c)
			answer := evaluate(&stack)
			stack.Push(answer)
		}
	}
	value, _ := strconv.Atoi(stack.Pop())
	return value
}

func evaluate(stack *Stack) string {
	op1, _ := strconv.Atoi(stack.Pop())
	operator := stack.Pop()
	op2, _ := strconv.Atoi(stack.Pop())
	switch operator {
	case "+":
		return strconv.Itoa(op1 + op2)
	case "*":
		return strconv.Itoa(op1 * op2)
	default:
		fmt.Println("Unknown operator:", operator)
		return "-1"
	}
}
