package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sample = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func main() {
	// part 1
	fmt.Printf("Invalid sample sum total: %d\n", parseForInvalidNums(sample, isInvalid))
	fmt.Printf("Invalid data   sum total: %d\n", parseForInvalidNums(load(), isInvalid))

	// part 2
	fmt.Printf("Invalid sample sum total: %d\n", parseForInvalidNums(sample, isInvalidPart2))
	fmt.Printf("Invalid data   sum total: %d\n", parseForInvalidNums(load(), isInvalidPart2))
}

func parseForInvalidNums(input string, isInvalidFunc func(int) bool) int {
	var sum int
	ranges := strings.Split(strings.TrimSuffix(input, "\n"), ",")
	for _, r := range ranges {
		thisRange := strings.Split(r, "-")
		first, _ := strconv.Atoi(thisRange[0])
		second, _ := strconv.Atoi(thisRange[1])
		for i := first; i <= second; i++ {
			if isInvalidFunc(i) {
				sum += i
			}
		}
	}
	return sum
}

func isInvalidPart2(num int) bool {
	var numInRunes = []rune(strconv.Itoa(num))
	if len(numInRunes) < 2 {
		return false
	}
	
	curr := ""
	for i := 0; i < len(numInRunes); i++ {
		curr += string(numInRunes[i])
		repeatVal := len(numInRunes) / len(curr)
		if repeatVal > 1 {
			if string(numInRunes) == strings.Repeat(curr, repeatVal) {
				//fmt.Printf("Found duplicate: %s\n", string(numInRunes))
				return true
			}

		}
	}
	return false
}

func isInvalid(num int) bool {
	// second implementation
	var numInRunes = []rune(strconv.Itoa(num))
	if len(numInRunes) <= 1 || len(numInRunes)%2 != 0 {
		return false
	}

	// check if even length of string
	half := len(numInRunes) / 2

	if string(numInRunes[:half]) == string(numInRunes[half:]) {
		return true
	}
	return false

	// original implementation, works for example
	//for i := 0; i < len(numInRunes); i++ {
	//	curr := numInRunes[:i+1]
	//	potentialDuplicate := numInRunes[i+1:]
	//	if string(curr) == string(potentialDuplicate) {
	//		return true
	//	}
	//}
	//return false
}

func load() string {
	f, readErr := os.ReadFile("data.txt")
	if readErr != nil {
		panic(readErr)
	}
	return string(f)
}
