package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Validator struct {
	min int
	max int
}

func (v Validator) Validate(value int) bool {
	return value >= v.min && value <= v.max
}

func (v Validator) String() string {
	return fmt.Sprintf("min: %d, max: %d", v.min, v.max)
}

func NewValidator(line string) *Validator {
	values := strings.Split(line, "-")
	v := Validator{}
	v.min, _ = strconv.Atoi(values[0])
	v.max, _ = strconv.Atoi(values[1])

	return &v
}

type CompositeValidator struct {
	name       string
	validators []*Validator
}

func (c CompositeValidator) Validate(value int) bool {
	for _, validator := range c.validators {
		if validator.Validate(value) {
			return true
		}
	}
	return false
}

func (c CompositeValidator) String() string {
	return fmt.Sprintf(c.name)
}

func NewCompositeValidator(line string) *CompositeValidator {
	result := strings.Split(line, ": ")
	c := CompositeValidator{}
	c.name = result[0]
	for _, v := range strings.Split(result[1], " or ") {
		c.validators = append(c.validators, NewValidator(v))
	}

	return &c
}

type Ticket struct {
	values []int
}

func (t Ticket) String() string {
	return fmt.Sprintf("%v", t.values)
}

func NewTicket(line string) *Ticket {
	ticket := Ticket{}
	values := strings.Split(line, ",")
	for _, value := range values {
		v, _ := strconv.Atoi(value)
		ticket.values = append(ticket.values, v)
	}

	return &ticket
}

type Guess struct {
	fieldName    string
	fieldIndexes []int
}

func (g Guess) String() string {
	return fmt.Sprintf("%s: %v", g.fieldName, g.fieldIndexes)
}

type Guesses []*Guess

func (g Guesses) Len() int {
	return len(g)
}

func (g Guesses) Less(i, j int) bool {
	return len(g[i].fieldIndexes) < len(g[j].fieldIndexes)
}

func (g Guesses) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func main() {
	scanner := getScanner()
	validators := getValidators(scanner)
	myTicket := getMyTicket(scanner)
	nearbyTickets := getNearbyTickets(scanner)

	validTickets := []*Ticket{}
	invalidValues := 0
	for _, ticket := range nearbyTickets {
		ok, invalidValue := isTicketValid(ticket, validators)
		if ok {
			validTickets = append(validTickets, ticket)
		} else {
			invalidValues += invalidValue
		}
	}
	// Part A
	fmt.Println(invalidValues)
	guesses := getGuesses(validTickets, validators)
	fieldMap := getFieldMap(guesses)
	decodeTicket(myTicket, fieldMap)
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

func getValidators(scanner *bufio.Scanner) []*CompositeValidator {
	validators := []*CompositeValidator{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		validators = append(validators, NewCompositeValidator(line))
	}
	return validators
}

func getMyTicket(scanner *bufio.Scanner) *Ticket {
	scanner.Scan()
	scanner.Scan()

	ticket := NewTicket(scanner.Text())
	return ticket
}

func getNearbyTickets(scanner *bufio.Scanner) []*Ticket {
	tickets := []*Ticket{}
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		ticket := NewTicket(scanner.Text())
		tickets = append(tickets, ticket)
	}
	return tickets
}

/**
 * For a ticket to be valid, all fields must be valid.
 */
func isTicketValid(ticket *Ticket, validators []*CompositeValidator) (bool, int) {
	for _, value := range ticket.values {
		if !isFieldValid(value, validators) {
			return false, value
		}
	}
	return true, 0
}

/**
 * For a field to be valid, it must pass at least one validator
 */
func isFieldValid(value int, validators []*CompositeValidator) bool {
	for _, validator := range validators {
		if validator.Validate(value) {
			return true
		}
	}
	return false
}

func getGuesses(tickets []*Ticket, validators []*CompositeValidator) Guesses {
	guesses := Guesses{}
	for _, validator := range validators {
		guess := getGuess(tickets, validator)
		guesses = append(guesses, guess)
	}
	sort.Sort(guesses)
	return guesses
}

/*
 * This will return the index of the field that works for all validators.
 */
func getGuess(tickets []*Ticket, validator *CompositeValidator) *Guess {
	fields := make(map[int]bool)
	for _, ticket := range tickets {
		for i, value := range ticket.values {
			isValid := validator.Validate(value)
			index := i + 1
			if _, ok := fields[index]; !ok {
				fields[index] = isValid
			} else {
				fields[index] = fields[index] && isValid
			}
		}
	}
	validFields := []int{}
	for fieldIndex, isValid := range fields {
		if isValid {
			validFields = append(validFields, fieldIndex)
		}
	}
	sort.Ints(validFields)
	guess := Guess{}
	guess.fieldName = validator.name
	guess.fieldIndexes = validFields
	return &guess
}

/**
 * Returns a map of the field index to the field name to help decode the ticket values.
 */
func getFieldMap(guesses Guesses) map[int]string {
	fieldMap := make(map[int]string)
	for _, guess := range guesses {
		for _, index := range guess.fieldIndexes {
			if _, ok := fieldMap[index]; !ok {
				fieldMap[index] = guess.fieldName
				break
			}
		}
	}
	return fieldMap
}

/**
 * Given the map of the field index to the field name, this will decode the ticket
 * and solve the problem.
 */
func decodeTicket(ticket *Ticket, fieldMap map[int]string) {
	decodedTicket := make(map[string]int)
	partBSolve := 1
	for fieldIndex, fieldName := range fieldMap {
		decodedTicket[fieldName] = ticket.values[fieldIndex-1]
		if strings.HasPrefix(fieldName, "departure") {
			partBSolve *= decodedTicket[fieldName]
		}
	}
	fmt.Println(partBSolve)
}
