package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type List struct {
	arr []string
}

func (s *List) Push(e string) {
	s.arr = append(s.arr, e)
}

func (s *List) Pop() string {
	n := len(s.arr) - 1
	if n < 0 {
		panic("No elements.")
	}
	e := s.arr[n]
	s.arr = s.arr[:n]
	return e
}

func (s *List) Dequeue() string {
	n := len(s.arr)
	if n == 0 {
		panic("No elements")
	}
	e := s.arr[0]
	s.arr = s.arr[1:]
	return e
}

func (s *List) Peek() string {
	n := len(s.arr) - 1
	if n < 0 {
		panic("No elements.")
	}
	return s.arr[n]
}

func (s *List) IsEmpty() bool {
	return len(s.arr) == 0
}

func main() {
	scanner := getScanner()
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		value := evaluateExpression(line)
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
	input := strings.Split(line, " ")
	outputQueue := List{}
	operatorStack := List{}

	for _, c := range input {
		if c == " " {
			continue
		} else if c == "+" || c == "*" {
			if !operatorStack.IsEmpty() && (operatorStack.Peek() == "+") {
				operator := operatorStack.Pop()
				outputQueue.Push(operator)
			}
			operatorStack.Push(c)
		} else if c == "(" {
			operatorStack.Push(c)
		} else if c == ")" {
			for operatorStack.Peek() != "(" {
				operator := operatorStack.Pop()
				outputQueue.Push(operator)
			}
			operatorStack.Pop() // Remove the paranthesis
		} else {
			outputQueue.Push(c)
		}
	}
	for !operatorStack.IsEmpty() {
		operator := operatorStack.Pop()
		outputQueue.Push(operator)
	}
	return evaluate(&outputQueue)
}

func evaluate(queue *List) int {
	stack := List{}
	for !queue.IsEmpty() {
		value := queue.Dequeue()
		if value == "*" {
			operand1, _ := strconv.Atoi(stack.Pop())
			operand2, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(operand1 * operand2))
		} else if value == "+" {
			operand1, _ := strconv.Atoi(stack.Pop())
			operand2, _ := strconv.Atoi(stack.Pop())
			stack.Push(strconv.Itoa(operand1 + operand2))
		} else {
			stack.Push(value)
		}
	}
	value, _ := strconv.Atoi(stack.Pop())
	return value
}
