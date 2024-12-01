package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Erorr while reading the input file, ", err)
		return
	}
	defer file.Close()

	var left, right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), " ")
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

	resSum := 0
	similarityScoreMap := make(map[int]int)

	for idx := range left {
		similarityScoreMap[right[idx]]++
		leftNum := float64(left[idx])
		rightNum := float64(right[idx])
		diff := math.Abs(leftNum - rightNum)
		resSum += int(diff)
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
