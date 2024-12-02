package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isValidSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		diff := int(math.Abs(float64(numbers[i] - numbers[i-1])))

		if diff < 1 || diff > 3 {
			return false
		}

		if isIncreasing && numbers[i] <= numbers[i-1] {
			return false
		}
		if !isIncreasing && numbers[i] >= numbers[i-1] {
			return false
		}
	}
	return true
}

func canBeMadeValid(numbers []int) bool {
	// Try removing each number
	for i := range numbers {
		// Create new slice without number at index i
		newNumbers := make([]int, 0)
		newNumbers = append(newNumbers, numbers[:i]...)
		newNumbers = append(newNumbers, numbers[i+1:]...)

		if isValidSequence(newNumbers) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while reading input file:", err)
		return
	}
	defer file.Close()

	safeReports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		// Convert all numbers in the line
		numbers := make([]int, 0)
		for _, num := range line {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Failed to convert string to int: %s\n", num)
				continue
			}
			numbers = append(numbers, n)
		}

		if len(numbers) < 2 {
			continue
		}

		if isValidSequence(numbers) {
			safeReports++
			continue
		}

		// If not valid, check if it can be made valid by removing one number
		if canBeMadeValid(numbers) {
			safeReports++
		}
	}

	fmt.Println("Safe reports =", safeReports)
}
