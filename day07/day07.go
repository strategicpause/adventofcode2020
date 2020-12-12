package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	name      string
	children  map[string]int
	totalBags int
}

func main() {
	containsShinyGold := make(map[string]bool)
	bagMap := getBagMap()
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

func getBagMap() map[string]*Bag {
	childRegex, _ := regexp.Compile(CHILD_REGEX)
	bagMap := make(map[string]*Bag)
	scanner := getScanner()
	for scanner.Scan() {
		line := scanner.Text()
		bag := getBag(line, childRegex)
		bagMap[bag.name] = bag
	}
	return bagMap
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
	bag.children = make(map[string]int)
	bag.totalBags = -1
	str := strings.Split(line, SPLIT)
	bag.name = str[0]
	if str[1] == NO_BAGS {
		return &bag
	}
	str = strings.Split(str[1], COMMA_SEP)
	for _, s := range str {
		i := strings.Index(s, " ")
		num, _ := strconv.Atoi(s[:i])
		bagName := childRegex.FindStringSubmatch(s)[1]
		bag.children[bagName] = num
	}
	return &bag
}

func hasGoldBags(bag *Bag, bagMap map[string]*Bag, containsShinyGold map[string]bool) bool {
	if _, ok := containsShinyGold[bag.name]; ok {
		return containsShinyGold[bag.name]
	}
	atLeastOne := false
	for child, _ := range bag.children {
		if child == SHINY_GOLD {
			return true
		}
		atLeastOne = atLeastOne || hasGoldBags(bagMap[child], bagMap, containsShinyGold)
	}
	return atLeastOne
}

func getNumBags(bagName string, bagMap map[string]*Bag) int {
	total := 0
	bag := bagMap[bagName]
	if bag.totalBags >= 0 {
		return bag.totalBags
	}
	if len(bag.children) == 0 {
		bag.totalBags = 0
		return 0
	}

	for childBagName, numBags := range bag.children {
		childTotal := getNumBags(childBagName, bagMap)
		// +1 to include the bag itself
		total += (1 + childTotal) * numBags
	}
	bag.totalBags = total
	return total
}
