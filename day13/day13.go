package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	scanner := getScanner()
	scanner.Scan()
	time, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	lowest := 1000
	nextBus := 0
	for _, bus := range buses {
		if bus != "x" {
			n, _ := strconv.Atoi(bus)
			timeUntil := n - int(math.Mod(float64(time), float64(n)))
			if lowest > timeUntil {
				lowest = timeUntil
				nextBus = n
			}
		}
	}
	fmt.Println(lowest * nextBus)
}

type CRT struct {
	a *big.Int
	n *big.Int
}

/**
 * Used the Chinese Remainder Theorem (CRT) for this one as the brute force method
 * would take too long. In the time that it took to generate a solution, I
 * was able to learn how to use the CRT from https://rosettacode.org/wiki/Chinese_remainder_theorem
 * and apply it to my solution here.
 */
func part2() {
	scanner := getScanner()
	scanner.Scan()
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	crts := []*CRT{}
	for i, bus := range buses {
		if bus != "x" {
			crt := &CRT{}
			busNum, _ := strconv.Atoi(bus)
			crt.n = big.NewInt(int64(busNum))
			crt.a = big.NewInt(int64(busNum - i))
			crts = append(crts, crt)
		}
	}

	p := new(big.Int).Set(crts[0].n)
	for _, crt := range crts[1:] {
		p.Mul(p, crt.n)
	}

	var x, q, s, z big.Int
	for _, crt := range crts {
		q.Div(p, crt.n)
		z.GCD(nil, &s, crt.n, &q)
		x.Add(&x, s.Mul(crt.a, s.Mul(&s, &q)))
	}
	fmt.Println(x.Mod(&x, p))
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
