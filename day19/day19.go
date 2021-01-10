package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	Input    string
	Resolved string
}

func NewRule(line string) *Rule {
	rule := Rule{}
	rule.Input = line

	return &rule
}

func NewResolvedRule(line string) *Rule {
	rule := NewRule(line)
	rule.Resolved = line

	return rule
}

func main() {
	scanner := getScanner()
	// Indexed by rule number and rule
	rules := make(map[string]*Rule)
	runRegex := false
	var regex string
	numMatches := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			regex = GetRegex(rules)
			fmt.Println("Found Regex", regex)
			runRegex = true
		} else if runRegex {
			matched, _ := regexp.MatchString(regex, line)
			if matched {
				numMatches++
			}
		} else {
			i, rule := ParseRule(line)
			rules[i] = rule
		}
	}
	fmt.Println(numMatches)
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

func ParseRule(line string) (string, *Rule) {
	parts := strings.Split(line, ": ")
	index, rule := parts[0], parts[1]
	if strings.HasPrefix(rule, "\"") {
		return index, NewResolvedRule(string(rule[1]))
	} else {
		return index, NewRule(rule)
	}
}

// Answer inspired by https://github.com/mnml/aoc/blob/master/2020/19/2.go
func GetRegex(rules map[string]*Rule) string {
	rules["8"].Resolved = Resolve("42", rules) + "+"
	for i := 1; i <= 10; i++ {
		rules["11"].Resolved += fmt.Sprintf("|%s{%d}%s{%d}", Resolve("42", rules), i, Resolve("31", rules), i)
	}
	rules["11"].Resolved = "(" + rules["11"].Resolved[1:] + ")"
	return "^" + Resolve("0", rules) + "$"
}

func Resolve(index string, rules map[string]*Rule) (re string) {
	rule := rules[index]
	if len(rule.Resolved) > 0 {
		return rule.Resolved
	}
	for _, s := range strings.Split(rule.Input, " | ") {
		re += "|"
		for _, s := range strings.Fields(s) {
			re += Resolve(s, rules)
		}
	}
	rule.Resolved = "(" + re[1:] + ")"
	return rule.Resolved
}
