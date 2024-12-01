package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Erorr while reading the input file, ", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	similarityScoreMap := make(map[int]int)
	var left, right []int

	for _, line := range lines {
		lineSplitted := strings.Split(line, " ")
		fmt.Println(lineSplitted)

		leftNumb, err := strconv.Atoi(lineSplitted[0])
		if err != nil {
			fmt.Println("Erorr while converting string to int (left), ", err)
		}
		left = append(left, leftNumb)

		rightNumb, err2 := strconv.Atoi(lineSplitted[len(lineSplitted)-1])
		if err2 != nil {
			fmt.Println("Erorr while converting string to int (right), ", err)
		}
		right = append(right, rightNumb)
	}

	sort.Ints(left)
	sort.Ints(right)

	// Store left-side numbers to check the existence with the right side
	for _, num := range left {
		similarityScoreMap[num] = 0
	}

	resSum := 0
	idx := 0
	for idx < len(left) {
		_, exists := similarityScoreMap[right[idx]]
		if exists {
			similarityScoreMap[right[idx]]++
		}

		leftNum := float64(left[idx])
		rightNum := float64(right[idx])
		diff := math.Abs(leftNum - rightNum)
		resSum += int(diff)
		idx += 1
	}

	fmt.Println("resSum = ", resSum)
	// Correct answer = 2000468

	// PART 2 - calculate similarity score

	similarityScore := 0
	for key, val := range similarityScoreMap {
		similarityScore += key * val
	}
	fmt.Println("similarityScore = ", similarityScore)
	// Correct answer = 18567089
}
