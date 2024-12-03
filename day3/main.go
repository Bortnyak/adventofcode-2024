package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading input.txt")
		return
	}

	contentStr := string(content)
	matches := findMulPatterns(contentStr)
	result := 0
	flag := true

	for _, match := range matches {
		if match == "don't()" {
			flag = false
			continue
		}

		if match == "do()" {
			flag = true
			continue
		}

		if flag == false {
			continue
		}

		result += getA(match) * getB(match)
	}

	fmt.Println("result = ", result)
}

func findMulPatterns(input string) []string {
	pattern := `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return nil
	}

	matches := re.FindAllString(input, -1)

	return matches
}

func getA(match string) int {
	aStartIndex := strings.Index(match, "(") + 1
	aEndIndex := strings.Index(match, ",")
	a := match[aStartIndex:aEndIndex]

	aInt, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Failed to convert a to int")
	}

	return aInt
}

func getB(match string) int {
	bStartIndex := strings.Index(match, ",") + 1
	bEndIndex := strings.Index(match, ")")
	b := match[bStartIndex:bEndIndex]

	bInt, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Failed to parse b to int")
	}

	return bInt
}
