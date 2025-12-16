package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* 
With some help from https://pastebin.com/geca2G3G if it still exists
*/

var sample = `987654321111111
811111111111119
234234234234278
818181911112111`

func load() []string {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(b)), "\n")
}

func main() {
	var (
		//data = strings.Split(strings.TrimSpace(sample), "\n")
		data = load()
		sum  = 0
	)

	for _, d := range data {
		// Part 1
		sum += getMaxJoltage(d)

		// Part 2
		sum += getMaxJoltageOverride(d)
	}
	fmt.Printf("Sum of %d battery banks: %d\n", len(data), sum)
}

func getMaxJoltage(s string) int {
	var (
		r           = []rune(s)
		maxValIndex = 0
	)

	// find (first) max value in string (exclude last entry)
	for i := 1; i < len(r)-1; i++ {
		if r[i] > r[maxValIndex] {
			maxValIndex = i
		}
	}

	// find the second max value in the string beyond the first one found
	// - should be a subset of the overall string, and should only require one pass
	var otherMaxValIndex = maxValIndex + 1
	for i := otherMaxValIndex + 1; i < len(r); i++ {
		if r[i] > r[otherMaxValIndex] {
			otherMaxValIndex = i
		}
	}

	maxJoltage, _ := strconv.Atoi(string([]rune{r[maxValIndex], r[otherMaxValIndex]}))
	return maxJoltage
}
func getMaxJoltageOverride(s string) int {
	var (
		r         = []rune(s)
		batteries []rune
		start     = 0
		numNeeded = 12
	)

	// count down from what we know we need
	for numNeeded > 0 {
		var (
			maxVal rune
			maxIdx int
			end    = len(s) - numNeeded + 1
			// can effectively finish early if we need 5 chars and we have 5 left
		)

		// find max in the current set of batteries
		for i := start; i < end; i++ {
			// shortcut to get rune's int value if it's an int
			//v := int(s[i] - '0')
			if r[i] > maxVal {
				maxVal = r[i]
				maxIdx = i - start
			}
		}

		// append found max value to our set of batteries we need to turn on
		batteries = append(batteries, maxVal)
		start = start + maxIdx + 1
		numNeeded -= 1
	}

	jolts, _ := strconv.Atoi(string(batteries))
	return jolts
}
