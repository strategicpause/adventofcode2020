package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type HashSet struct {
	m map[int]bool
}

func NewSet() *HashSet {
	s := &HashSet{}
	s.m = make(map[int]bool)
	return s
}

func (s *HashSet) Add(value int) {
	s.m[value] = true
}

func (s *HashSet) Remove(value int) {
	s.m[value] = false
}

func (s *HashSet) Contains(value int) bool {
	_, c := s.m[value]
	return c
}

func (s *HashSet) ToList() *[]int {
	var list []int
	for key, value := range s.m {
		if value {
			list = append(list, key)
		}
	}
	return &list
}

func main() {
	a()
	b()
}

func a() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input.txt")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	set := NewSet()
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if set.Contains(2020 - n) {
			fmt.Printf("%d\n", n*(2020-n))
		}
		set.Add(n)
	}
}

func b() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading input.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	set := NewSet()
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		set.Add(n)
	}

	m := set.ToList()
	n := len(*m)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if set.Contains(2020 - (*m)[i] - (*m)[j]) {
				fmt.Printf("%d\n", (*m)[i]*(*m)[j]*(2020-(*m)[i]-(*m)[j]))
				return
			}
		}
	}
	fmt.Printf("Not Found.")
}
