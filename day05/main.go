package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func load(fileName string) []string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// trim out leading and trailing spaces
	return strings.Split(strings.TrimSpace(string(f)), "\n")
}

type Range struct {
	start int
	end   int
}

// Used for Part 1, builds a slice of Range structs that track the start and end points of each range entry,
// which allows us to more efficiently understand if a given ingredient is considered fresh.
func buildRanges(input []string) ([]Range, []int) {
	var (
		ingredientRanges []Range
		ingredients      []int
	)

	// iterate over the input
	for _, line := range input {
		// ingredientLookup come first
		if strings.Contains(line, "-") {
			var (
				spl       = strings.Split(line, "-")
				first, _  = strconv.Atoi(spl[0])
				second, _ = strconv.Atoi(spl[1])
			)
			ingredientRanges = append(ingredientRanges, Range{first, second})
		} else if line != " " {
			ing, _ := strconv.Atoi(strings.TrimSpace(line))
			ingredients = append(ingredients, ing)
		}
	}

	sort.Slice(ingredientRanges, func(i, j int) bool {
		return ingredientRanges[i].start < ingredientRanges[j].start
	})

	return ingredientRanges, ingredients
}

// Used for Part 1, determines the number of fresh ingredients found based on a set of precomputed Range entries.
func determineFreshIngredients(ingredientRanges []Range, ingredients []int) int {
	sum := 0
	for _, i := range ingredients {
		for _, j := range ingredientRanges {
			if i >= j.start && i <= j.end {
				sum += 1
				break
			}
		}
	}
	return sum
}

// Used for Part 2
func countFreshIngredientIds(input []string) int {
	var ingredientRanges []Range
	// add each range to our slice
	for _, line := range input {
		if strings.Contains(line, "-") {
			spl := strings.Split(line, "-")
			first, _ := strconv.Atoi(spl[0])
			second, _ := strconv.Atoi(spl[1])
			ingredientRanges = append(ingredientRanges, Range{first, second})
		}
	}

	// sort them based on their first entry
	sort.Slice(ingredientRanges, func(i, j int) bool {
		return ingredientRanges[i].start < ingredientRanges[j].start
	})

	// go through each range and start merging together ranges
	var mergedRanges []Range
	for {
		// no need to compare anymore, we've already compared them all
		if len(ingredientRanges) == 1 {
			mergedRanges = append(mergedRanges, ingredientRanges...)
			break
		}

		// pop off both to compare
		first := ingredientRanges[0]
		second := ingredientRanges[1]
		ingredientRanges = ingredientRanges[2:]

		// ranges do not intersect
		if second.start > first.end {
			// this is a valid range, add to our new queue
			mergedRanges = append(mergedRanges, first)
			// add the second item back into the front of the queue
			ingredientRanges = append([]Range{second}, ingredientRanges...)
		} else {
			// ranges do intersect!
			ingredientRanges = append([]Range{{first.start, max(first.end, second.end)}}, ingredientRanges...)
		}
	}

	// finally, count up numerical sums
	sum := 0
	for _, r := range mergedRanges {
		sum += r.end - r.start + 1 // inclusive
	}
	return sum
}

func main() {
	// Part 1
	currTime := time.Now()
	ranges, ingredients := buildRanges(load("sample.txt"))
	fmt.Printf("Found %d fresh ingredients in sample (%s)\n",
		determineFreshIngredients(ranges, ingredients), time.Since(currTime))

	// Part 1
	currTime = time.Now()
	ranges, ingredients = buildRanges(load("data.txt"))
	fmt.Printf("Found %d fresh ingredients in data (%s)\n",
		determineFreshIngredients(ranges, ingredients), time.Since(currTime))

	// Part 2
	currTime = time.Now()
	fmt.Printf("Num unique fresh ingredient IDs in sample: %d (%s)\n",
		countFreshIngredientIds(load("sample.txt")), time.Since(currTime))

	currTime = time.Now()
	fmt.Printf("Num unique fresh ingredient IDs in data: %d (%s)\n",
		countFreshIngredientIds(load("data.txt")), time.Since(currTime))
}
