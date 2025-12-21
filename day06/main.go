package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Part 1
	partOneSampleValues, partOneSampleInstructions := part1Load("sample.txt")
	currTime := time.Now()
	fmt.Printf("Final Total Sum (sample.txt): %d (%s)\n",
		part1Calculate(partOneSampleValues, partOneSampleInstructions), time.Since(currTime))

	// Part 1
	partOneDataValues, partOneDataInstructions := part1Load("data.txt")
	currTime = time.Now()
	fmt.Printf("Final Total Sum (data.txt)  : %d (%s)\n",
		part1Calculate(partOneDataValues, partOneDataInstructions), time.Since(currTime))

	// Part 2
	currTime = time.Now()
	fmt.Printf("Final Total Sum (sample.txt): %d (%s)\n", part2("sample.txt"), time.Since(currTime))
	currTime = time.Now()
	fmt.Printf("Final Total Sum (data.txt)  : %d (%s)\n", part2("data.txt"), time.Since(currTime))
}

// Just loads the original file and returns a grid of integers and the []string of operators.
func part1Load(filename string) ([][]int, []string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var (
		lines        = strings.Split(strings.TrimSpace(string(f)), "\n")
		cols         = 0
		allNums      [][]int
		instructions []string
	)

	for _, l := range lines {
		var nums []int
		elements := strings.Split(l, " ")
		for _, e := range elements {
			if e != "" {
				// final row needs to remain strings
				if e == "*" || e == "+" {
					instructions = append(instructions, e)
					continue
				}
				// every other row needs to be converted and appended
				eInt, _ := strconv.Atoi(e)
				nums = append(nums, eInt)
			}
		}

		// on the last row, nums will be nil as we didn't add to it, so we only need
		// to check for nil here. Otherwise we'd need to loop through the input twice,
		// or rely on the last row in its own explicit operation (also possible).
		if nums != nil {
			// track this on first row only
			if cols == 0 {
				cols = len(nums)
			}
			// later validation
			if len(nums) != cols {
				panic(fmt.Sprintf("Received %d nums, expected %d for line: %s", len(nums), cols, l))
			}
			allNums = append(allNums, nums)
		}

	}
	return allNums, instructions
}

// Used for Part 1 only, calculates the final sum from the previously loaded integer grid
// and operator instruction set.
func part1Calculate(values [][]int, instructions []string) int {
	sum := 0
	for col, ins := range instructions {
		// collect numbers for this particular column
		var thisCol []int
		for row := 0; row < len(values); row++ {
			thisCol = append(thisCol, values[row][col])
		}

		// calculate total for this column
		switch ins {
		case "*":
			colProduct := thisCol[0]
			for c := 1; c < len(thisCol); c++ {
				colProduct = colProduct * thisCol[c]
			}
			sum += colProduct
		case "+":
			colSum := thisCol[0]
			for c := 1; c < len(thisCol); c++ {
				colSum += thisCol[c]
			}
			sum += colSum
		}
	}
	return sum
}

// Part 2, both loads the file in question and processes its sum based on the more complicated summation
// logic as specified. This requires a lot of transposition of values and breaking down characters and
// then rejoining them together, only converting them to integers and adding them at the end.
func part2(filename string) int {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var (
		lines           = strings.Split(strings.TrimSuffix(string(f), "\n"), "\n")
		operators       = lines[len(lines)-1]
		operatorIndexes []int
	)
	// pop off last row of operators so that we can handle them separately
	lines = lines[:len(lines)-1]

	// find starting indexes for each column of number representations
	// so for:
	// 123 456
	// +   *
	// ^...^...
	// ...would be: []int{0, 4}
	for i, o := range []rune(operators) {
		if string(o) == "*" || string(o) == "+" {
			operatorIndexes = append(operatorIndexes, i)
		}
	}

	// iterate right to left, from len(line)-1 to reversed index val
	slices.Reverse(operatorIndexes)

	// build up a strNumGrid of strings to match input
	var strNumGrid [][]string
	for _, line := range lines {
		var strNumsWithSpaces []string
		for _, idx := range operatorIndexes {
			// utilize the index values from the operators slice to additionally
			// slice the strNumsWithSpaces 2d slice. This will build string representations
			// of each number of each "column" of each row.
			strNumsWithSpaces = append(strNumsWithSpaces, line[idx:])

			// trim off the last entry of the line to continue iterating from whatever is "last"
			if idx != 0 {
				line = line[:idx-1]
			}
		}
		// Since we iterated right to left, we can reverse the strings we found to put them
		// back into the original order (makes comparisons easier, not technically needed).
		slices.Reverse(strNumsWithSpaces)
		strNumGrid = append(strNumGrid, strNumsWithSpaces)
	}

	// reverse again the operatorIndexes to better handle this next part
	slices.Reverse(operatorIndexes)

	// this is the final value, added from the respective +/* operations of each column
	totalSumOfAllCols := 0

	// iterate over each "column" of found strings in the strNumGrid we have built, iterating
	// down each column from left to right.
	for y := 0; y < len(strNumGrid[0]); y++ {
		// build up a double slice that represents a grid of string characters to put this whole
		// column of numbers from the grid into its own "sub grid" so that we can track each position
		// of numbers to combine into final, vertically placed numbers.
		var col = make([][]string, len(strNumGrid[0][y]))
		for x := 0; x < len(strNumGrid); x++ {
			// We need to break down each number into a slice containing all its characters
			// "123" becomes []string{"1", "2", "3"}
			// ...which we can then place into its own slot in the [][]string 2d "total" slice
			var str = strNumGrid[x][y]
			for i := 0; i < len(str); i++ {
				col[i] = append(col[i], string(str[i]))
			}
			//fmt.Println(col)
		}
		// At this point we effectively have:
		// [[1] [2] [3]]
		// [[ ] [4] [5]]
		// [[ ] [ ] [6]]
		// ...for this particular column we're working on
		colSum := 0
		for transposedCol := 0; transposedCol < len(col); transposedCol++ {
			// now that we've swapped numbers around from column to row ordering, we can join them
			// into the final string that was vertical for this column of numbers, taking note
			// to remove the extra spaces that we no longer need, since we're now joining numbers
			// that were once vertically placed.
			itemColNum, _ := strconv.Atoi(strings.TrimSpace(strings.Join(col[transposedCol], "")))

			// switch on operators[operatorIndexes[y]] because this column we're currently processing is the same
			// index as being used in operators, and operatorIndexes contains the index values for the appropriate
			// operators slice - this multiplies or adds numbers as intended.
			switch string(operators[operatorIndexes[y]]) {
			case "*":
				if colSum == 0 {
					colSum = itemColNum
				} else {
					colSum *= itemColNum
				}
			case "+":
				colSum += itemColNum
			}
		}
		// we've determined a column number, add it to the total that we will sum up across all columns.
		totalSumOfAllCols += colSum
	}
	return totalSumOfAllCols
}
