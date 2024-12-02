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

		// Convert all numbers in the line upfront
		numbers := make([]int, len(line))
		for i, num := range line {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Failed to convert string to int: %s\n", num)
				continue
			}
			numbers[i] = n
		}

		if len(numbers) < 2 {
			continue
		}

		// Check if sequence is valid
		isValid := true
		isIncreasing := numbers[1] > numbers[0]

		for i := 1; i < len(numbers); i++ {
			diff := int(math.Abs(float64(numbers[i] - numbers[i-1])))

			// Check if difference is between 1 and 3
			if diff < 1 || diff > 3 {
				isValid = false
				break
			}

			// Check if sequence maintains direction (increasing or decreasing)
			if isIncreasing && numbers[i] <= numbers[i-1] {
				isValid = false
				break
			}
			if !isIncreasing && numbers[i] >= numbers[i-1] {
				isValid = false
				break
			}
		}

		if isValid {
			safeReports++
		}
	}

	fmt.Println("Safe reports =", safeReports)
}
