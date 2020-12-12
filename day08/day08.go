package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NOP   = "nop"
	ACC   = "acc"
	JMP   = "jmp"
	DEBUG = false
)

type Instruction struct {
	operator string
	operand  int
}

type ExecutionEnvironment struct {
	accumulator        int
	instructionPointer int
}

func main() {
	instructions := getInstructions()
	n := len(instructions)
	for i := 0; i < n; i++ {
		permute(instructions[i])
		hasLoop(instructions)
		permute(instructions[i])
	}
}

func getInstructions() []*Instruction {
	scanner := getScanner()
	instructions := []*Instruction{}
	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
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

func parseInstruction(line string) *Instruction {
	instruction := Instruction{}

	s := strings.Split(line, " ")
	instruction.operator = s[0]
	instruction.operand, _ = strconv.Atoi(s[1][1:])
	if s[1][0] == '-' {
		instruction.operand *= -1
	}

	return &instruction
}

func hasLoop(instructions []*Instruction) bool {
	ctx := ExecutionEnvironment{}
	numInstructions := len(instructions)
	instructionsSeen := map[int]bool{}

	for ctx.instructionPointer < numInstructions {
		ptr := ctx.instructionPointer
		if instructionsSeen[ptr] {
			return true
		}
		instructionsSeen[ptr] = true
		instruction := instructions[ptr]
		switch instruction.operator {
		case JMP:
			ctx.instructionPointer += instruction.operand
			debug("IP set to ", ctx.instructionPointer)
		case ACC:
			ctx.accumulator += instruction.operand
			ctx.instructionPointer++
			debug("ACC set to ", ctx.accumulator)
		case NOP:
			ctx.instructionPointer++
		}
	}
	fmt.Println(ctx.accumulator)
	return false
}

func debug(a ...interface{}) {
	if DEBUG {
		fmt.Println(a...)
	}
}

func permute(instruction *Instruction) {
	operator := instruction.operator
	if operator == NOP {
		instruction.operator = JMP
	} else if operator == JMP {
		instruction.operator = NOP
	}
}
