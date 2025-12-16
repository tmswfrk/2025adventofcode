package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func loadSample(filename string) [][]string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var (
		rows = strings.Split(strings.TrimSpace(string(b)), "\n")
		data [][]string
	)
	for i := 0; i < len(rows); i++ {
		data = append(data, strings.Split(strings.TrimSpace(rows[i]), ""))
	}
	return data
}

func main() {
	// Sample data, Part 1
	sampleTxt := "sample.txt"
	currTime := time.Now()
	fmt.Printf("%s: Total Papers Found  : %d (%s)\n", sampleTxt,
		discoverPaper(sampleTxt, false), time.Since(currTime))

	// Given data, Part 1
	data := "data.txt"
	currTime = time.Now()
	fmt.Printf("%s  : Total Papers Found  : %d (%s)\n", data,
		discoverPaper(data, false), time.Since(currTime))

	// Sample data, Part 2
	currTime = time.Now()
	fmt.Printf("%s: Total Papers Removed: %d (%s)\n", sampleTxt,
		repeatableDiscoverPaper(loadSample(sampleTxt), 0, false), time.Since(currTime))

	// Given data, Part 2
	currTime = time.Now()
	fmt.Printf("%s  : Total Papers Removed: %d (%s)\n", data,
		repeatableDiscoverPaper(loadSample(data), 0, false), time.Since(currTime))
}

// Used for part 2
func repeatableDiscoverPaper(data [][]string, totalPapersRemoved int, toPrint bool) int {
	var (
		papersRemoved = 0
		afterRemoval  [][]string
	)
	for i := 0; i < len(data); i++ {
		// set up a new row that will contain the changed state of the current data[i] row
		// this will be used to build up a similar [][]string that will be used as input
		// into another invocation of this same function.
		var thisRowFinal []string
		for j := 0; j < len(data[0]); j++ {
			// we only care to proceed if this immediate position is a paper
			if data[i][j] != "@" {
				thisRowFinal = append(thisRowFinal, ".")
				continue
			}

			// track a sum of how many neighbors of paper that this paper has
			numPapersNearby := 0

			// check in each direction for another roll of paper
			xDir, yDir := []int{-1, 0, 1}, []int{-1, 0, 1}
			for _, x := range xDir {
				for _, y := range yDir {
					if x == 0 && y == 0 {
						continue // skip, this is the item itself that we know is already a paper
					}
					toCheckI := i + x
					toCheckJ := j + y
					// boundary check
					if toCheckI < 0 || toCheckI >= len(data) || toCheckJ < 0 || toCheckJ >= len(data[0]) {
						continue
					}

					if data[toCheckI][toCheckJ] == "@" {
						numPapersNearby += 1
					}
				}
			}

			// if we have 4 or fewer papers nearby, we can remove this one.
			if numPapersNearby < 4 {
				thisRowFinal = append(thisRowFinal, ".")
				papersRemoved += 1
			} else {
				// track this for later, as this paper is NOT removed
				thisRowFinal = append(thisRowFinal, "@")
			}
		}
		// append the final row representation to the newly created [][]string
		// - could probably also swap out the existing data [][]string to save allocations
		afterRemoval = append(afterRemoval, thisRowFinal)
	}
	// Use this to track the final total, passed down and added to within each function invocation
	totalPapersRemoved += papersRemoved

	if toPrint {
		for _, r := range afterRemoval {
			fmt.Println(r)
		}
		fmt.Printf("Papers Removed: %d\n", papersRemoved)
		fmt.Printf("Total Papers Removed so Far: %d\n\n", totalPapersRemoved)
	}

	// base case
	if papersRemoved == 0 {
		return totalPapersRemoved
	}
	// Recursion!
	return repeatableDiscoverPaper(afterRemoval, totalPapersRemoved, toPrint)
}

// Used for part 1, slightly different function signature than part 2
func discoverPaper(source string, toPrint bool) int {
	data := loadSample(source)
	totalPapers := 0

	// keeping track of the next state of the room was a good set up for Part 2
	var final [][]string
	for i := 0; i < len(data); i++ {
		var thisRowFinal []string
		for j := 0; j < len(data[0]); j++ {
			// we only care to proceed if this immediate position is a paper
			if data[i][j] != "@" {
				thisRowFinal = append(thisRowFinal, ".")
				continue
			}

			// track a sum of how many neighbors of paper that this paper has
			numPapersNearby := 0

			// check in each direction for another roll of paper
			xDir, yDir := []int{-1, 0, 1}, []int{-1, 0, 1}
			for _, x := range xDir {
				for _, y := range yDir {
					if x == 0 && y == 0 {
						continue // skip, this is the item itself that we know is already a paper
					}
					toCheckI := i + x
					toCheckJ := j + y
					if toCheckI < 0 || toCheckI >= len(data) || toCheckJ < 0 || toCheckJ >= len(data[0]) {
						continue
					}

					if data[toCheckI][toCheckJ] == "@" {
						numPapersNearby += 1
					}
				}
			}

			if numPapersNearby < 4 {
				thisRowFinal = append(thisRowFinal, "x")
				totalPapers += 1
			} else {
				thisRowFinal = append(thisRowFinal, "@")
			}
		}
		final = append(final, thisRowFinal)
	}

	if toPrint {
		for _, r := range final {
			fmt.Println(r)
		}
		fmt.Printf("Total papers that can be removed: %d\n\n", totalPapers)
	}
	return totalPapers
}
