package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Direction struct {
	dx, dy int
}

func findXMAS(input string) int {
	// Split the input into lines to create the grid
	grid := strings.Split(strings.TrimSpace(input), "\n")

	rows := len(grid)
	cols := len(grid[0])
	count := 0
	foundPositions := make([]string, 0)

	// All possible directions
	directions := []Direction{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	isValid := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	checkDirection := func(row, col int, dir Direction) bool {
		if !isValid(row, col) {
			return false
		}

		// Check if we can go 4 positions in this direction
		if !isValid(row+3*dir.dx, col+3*dir.dy) {
			return false
		}

		var word strings.Builder
		word.Grow(4)
		for i := 0; i < 4; i++ {
			word.WriteByte(grid[row+i*dir.dx][col+i*dir.dy])
		}

		wordStr := word.String()
		found := wordStr == "XMAS"
		if found {
			foundPositions = append(foundPositions, fmt.Sprintf("(%d,%d) -> (%d,%d): %s",
				row, col, dir.dx, dir.dy, wordStr))
		}
		return found
	}

	// Check each position in the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				if checkDirection(i, j, dir) {
					count++
				}
			}
		}
	}

	// Debug output
	fmt.Println("Grid dimensions:", rows, "x", cols)
	fmt.Println("Sample found positions (first 5):")
	for i := 0; i < int(math.Min(float64(5), float64(len(foundPositions)))); i++ {
		fmt.Println(foundPositions[i])
	}

	return count
}

func main() {
	// Test with the example grid
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading input.txt")
		return
	}

	result := findXMAS(string(content))
	fmt.Printf("Found %d instances of XMAS\n", result) // Should print 18
}
