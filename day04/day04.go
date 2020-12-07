package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport map[string]string

func main() {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	passports := parsePassports()
	valid := 0
	for _, passport := range passports {
		validPassport := validatePassport(passport, fields)
		if validPassport {
			valid++
		}
	}
	fmt.Println(valid)
}

func parsePassports() []*Passport {
	passports := []*Passport{}

	lines := []string{}
	scanner := getScanner()
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passport := parsePassport(lines)
			passports = append(passports, passport)
			lines = []string{}
		} else {
			lines = append(lines, text)
		}
	}

	return passports
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

func parsePassport(passportLines []string) *Passport {
	passport := make(Passport)
	for _, line := range passportLines {
		fields := strings.Split(line, " ")
		for _, field := range fields {
			s := strings.Split(field, ":")
			passport[s[0]] = s[1]
		}
	}
	return &passport
}

func validatePassport(passport *Passport, fields []string) bool {
	for _, field := range fields {
		if val, ok := (*passport)[field]; !ok {
			return false
		} else {
			switch field {
			case "byr":
				byr, _ := strconv.Atoi(val)
				if byr < 1920 || byr > 2002 {
					return false
				}
			case "iyr":
				iyr, _ := strconv.Atoi(val)
				if iyr < 2010 || iyr > 2020 {
					return false
				}
			case "eyr":
				eyr, _ := strconv.Atoi(val)
				if eyr < 2020 || eyr > 2030 {
					return false
				}
			case "hgt":
				found, _ := regexp.MatchString("[0-9]{2,3}(in|cm)", val)
				if !found {
					return false
				}
				n := len(val)
				height, _ := strconv.Atoi(val[:n-2])
				unit := val[n-2:]
				if unit == "in" && (height < 59 || height > 76) {
					return false
				} else if unit == "cm" && (height < 150 || height > 193) {
					return false
				}
			case "hcl":
				found, _ := regexp.MatchString("#[0-9a-f]{6}", val)
				if !found {
					return false
				}
			case "ecl":
				found, _ := regexp.MatchString("(amb|blu|brn|gry|grn|hzl|oth)", val)
				if !found {
					return false
				}
			case "pid":
				found, _ := regexp.MatchString("[0-9]{9}", val)
				if !found {
					return false
				}
			default:
				return false
			}
		}
	}
	return true
}
