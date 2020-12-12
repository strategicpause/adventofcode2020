package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	SPLIT       = " bags contain "
	NO_BAGS     = "no other bags."
	COMMA_SEP   = ", "
	CHILD_REGEX = "\\d (.*) (bag|bags)\\.?"
	SHINY_GOLD  = "shiny gold"
)

type Bag struct {
	name     string
	children []string
}

func main() {
	childRegex, _ := regexp.Compile(CHILD_REGEX)
	containsShinyGold := make(map[string]bool)
	bagMap := make(map[string]*Bag)
	scanner := getScanner()
	for scanner.Scan() {
		line := scanner.Text()
		bag := getBag(line, childRegex)
		bagMap[bag.name] = bag
	}

	for bagName, bag := range bagMap {
		containsShinyGold[bagName] = hasGoldBags(bag, bagMap, containsShinyGold)
	}
	containsGold := 0
	for _, value := range containsShinyGold {
		if value {
			containsGold++
		}
	}
	fmt.Println(containsGold)
	numBags := getNumBags(SHINY_GOLD, bagMap)
	fmt.Println(numBags)
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

func getBag(line string, childRegex *regexp.Regexp) *Bag {
	bag := Bag{}
	str := strings.Split(line, SPLIT)
	bag.name = str[0]
	if str[1] == NO_BAGS {
		return &bag
	}
	str = strings.Split(str[1], COMMA_SEP)
	for _, s := range str {
		bagName := childRegex.FindStringSubmatch(s)[1]
		bag.children = append(bag.children, bagName)
	}
	return &bag
}

func hasGoldBags(bag *Bag, bagMap map[string]*Bag, containsShinyGold map[string]bool) bool {
	if _, ok := containsShinyGold[bag.name]; ok {
		return containsShinyGold[bag.name]
	}
	atLeastOne := false
	for _, child := range bag.children {
		if child == SHINY_GOLD {
			return true
		}
		atLeastOne = atLeastOne || hasGoldBags(bagMap[child], bagMap, containsShinyGold)
	}
	return atLeastOne
}

func getNumBags(bagName string, bagMap map[string]*Bag) {

}
