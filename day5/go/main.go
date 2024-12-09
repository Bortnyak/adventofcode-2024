package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read file
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file content
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Initialize maps and slices
	afterNum := make(map[string][]string)
	var resultListToCheckMid [][]string

	// Process each line
	for _, elem := range lines {
		elemLen := len(elem)

		if elemLen == 5 {
			left := elem[:2]
			right := elem[3:]

			if list, exists := afterNum[left]; exists {
				afterNum[left] = append(list, right)
			} else {
				afterNum[left] = []string{right}
			}
		} else if elemLen > 5 {
			elementsList := strings.Split(elem, ",")
			elements := make([]string, len(elementsList))
			copy(elements, elementsList)

			// Reverse elements
			for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
				elements[i], elements[j] = elements[j], elements[i]
			}

			isIntersection := false

			for index := range elements {
				elemToCheck := elements[index]
				listToCheckIn := elements[index+1:]

				if elemList, exists := afterNum[elemToCheck]; exists {
					intersectionList := intersection(listToCheckIn, elemList)
					if len(intersectionList) > 0 {
						isIntersection = true
						break
					}
				}
			}

			if !isIntersection {
				resultListToCheckMid = append(resultListToCheckMid, elementsList)
			}
		}
	}

	fmt.Println("result_list_to_check_mid =", resultListToCheckMid)

	// Calculate sum
	sum := 0
	if len(resultListToCheckMid) > 0 {
		for _, l := range resultListToCheckMid {
			elemListLen := len(l)
			middleIndex := (elemListLen - 1) / 2
			middleElem, err := strconv.Atoi(strings.TrimSpace(l[middleIndex]))
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				continue
			}
			sum += middleElem
		}
	}

	fmt.Println("sum =", sum)
}

// intersection returns the intersection of two string slices
func intersection(lst1, lst2 []string) []string {
	// Create a map to store elements from lst1
	set := make(map[string]bool)
	for _, item := range lst1 {
		set[item] = true
	}

	// Create result slice
	var result []string
	for _, item := range lst2 {
		if set[item] {
			result = append(result, item)
		}
	}

	return result
}
