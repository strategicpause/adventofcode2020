package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partA()
	partB()
}

func partA() {
	scanner := getScanner()
	oneMask := int64(0)
	zeroMask := int64(^0)
	memoryMap := make(map[int]int64)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			zeroMask, oneMask = parseMask(line)
		} else {
			address, value := parseMem(line)
			value |= oneMask
			value &= zeroMask
			memoryMap[address] = value
		}
	}
	sum := int64(0)
	for _, value := range memoryMap {
		sum += value
	}
	fmt.Println(sum)
}

func partB() {
	scanner := getScanner()
	memoryMap := make(map[int64]int64)
	mask := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask") {
			mask = getMask(line)
		} else {
			baseAddress, value := parseMem(line)
			addresses := getAddresses(baseAddress, mask)
			for _, address := range addresses {
				memoryMap[address] = value
			}
		}
	}
	sum := int64(0)
	for _, value := range memoryMap {
		sum += value
	}
	fmt.Println(sum)
}

func getScanner() *bufio.Scanner {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading sample.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return scanner
}

func parseMask(mask string) (int64, int64) {
	mask = getMask(mask)
	var oneMask int64 = 0
	var zeroMask int64 = ^0
	for i, m := range mask {
		if m != 'X' {
			val, _ := strconv.Atoi(string(m))
			if val == 1 {
				m := int64(1 << (35 - i))
				oneMask |= m
			} else if val == 0 {
				m := int64(^(1 << (35 - i)))
				zeroMask &= m
			}
		}
	}
	return zeroMask, oneMask
}

func getMask(line string) string {
	return strings.Split(line, "mask = ")[1]
}

func parseMem(mem string) (int, int64) {
	a := strings.Split(mem, " = ")
	addr := strings.Split(a[0], "mem[")[1]
	address, _ := strconv.Atoi(strings.TrimRight(addr, "]"))
	value, _ := strconv.Atoi(a[1])
	return address, int64(value)
}

func getAddresses(baseAddress int, mask string) []int64 {
	addresses := []int64{}
	address := fmt.Sprintf("%036b", baseAddress)
	var sb strings.Builder
	numX := 0
	for i, c := range mask {
		if c != '0' {
			sb.WriteRune(c)
		} else {
			sb.WriteByte(address[i])
		}
		if c == 'X' {
			numX++
		}
	}
	address = sb.String()
	n := int(math.Exp2(float64(numX)))
	for i := 0; i < n; i++ {
		sb.Reset()
		j := 0
		for _, c := range address {
			if c != 'X' {
				sb.WriteRune(c)
			} else if i&(1<<(j)) != 0 {
				sb.WriteString("1")
				j++
			} else {
				sb.WriteString("0")
				j++
			}
		}
		addressVal, _ := strconv.ParseInt(sb.String(), 2, 64)
		addresses = append(addresses, addressVal)
	}
	return addresses
}
