package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	value    int
	children []*Node
	total    int
}

func main() {
	list := getList()
	sort.Ints(list)
	solveA(list)
	solveB(list)
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

func solveA(list []int) {
	diffMap := map[int]int{}
	n := len(list)
	prev := list[0]
	// Diff from outlet
	diffMap[prev]++
	for i := 1; i < n; i++ {
		diff := list[i] - prev
		diffMap[diff]++
		prev = list[i]
	}
	// Diff to device
	diffMap[3]++
	fmt.Println(diffMap[1] * diffMap[3])
}

func solveB(list []int) {
	root := buildGraph(list)
	findAllPaths(root)
}

func buildGraph(list []int) *Node {
	graph := map[int]*Node{}
	// The outlet is represented by the root node.
	rootNode := newNode(0)
	graph[0] = rootNode
	for _, val := range list {
		node := newNode(val)
		graph[val] = node
		// Add the current node to nodes, 1, 2, and 3 values behind, if they exist.
		for i := 1; i < 4; i++ {
			prevNode := graph[val-i]
			if prevNode != nil {
				prevNode.children = append(prevNode.children, node)
			}
		}
	}
	// Return root node
	return rootNode
}

func newNode(val int) *Node {
	node := Node{}
	node.value = val
	node.children = []*Node{}
	node.total = math.MinInt32

	return &node
}

func findAllPaths(root *Node) {
	path := []int{}
	path = append(path, root.value)
	numPaths := findAllPaths2(root, path)
	fmt.Println(numPaths)
}

func findAllPaths2(root *Node, path []int) int {
	if root.total > 0 {
		return root.total
	}
	if len(root.children) == 0 {
		return 1
	}
	n := len(path)
	total := 0
	for _, child := range root.children {
		path := append(path, child.value)
		total += findAllPaths2(child, path)
		// Remove the last appended value
		path = path[:n]
	}
	// Memoize the total for each node
	root.total = total
	return total
}
