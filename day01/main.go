package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxDialVal = 100 // one more than max dial amount of 99 to count 0 as a valid entry
)

var sampleData = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func main() {
	data := loadData()
	fmt.Printf("Find  Zeros, Sample Data: %d\n", findZeros(50, sampleData))
	fmt.Printf("Count Zeros, Sample Data: %d\n", countZeros(50, sampleData))
	fmt.Printf("Find  Zeros, Full Data  : %d\n", findZeros(50, data))
	fmt.Printf("Count Zeros, Full Data  : %d\n", countZeros(50, data))
}

func splitInstruction(instruction string) (string, int) {
	amount, _ := strconv.Atoi(instruction[1:])
	return string(instruction[0]), amount
}

func findZeros(dial int, data []string) int {
	count := 0
	for _, ins := range data {
		dir, val := splitInstruction(ins)
		if dir == "L" {
			dial = dial - val
		} else {
			dial = dial + val
		}
		if dial%100 == 0 {
			count += 1
		}
	}
	return count
}

func countZeros(dial int, data []string) int {
	count := 0
	for _, ins := range data {
		dir, val := splitInstruction(ins)

		// iterate over each number
		for range val {
			if dir == "L" {
				dial = dial - 1
			} else {
				dial = dial + 1
			}
			// check if zero
			if dial%100 == 0 {
				count += 1
			}
		}
	}
	return count
}

func loadData() []string {
	var commands []string
	f, openErr := os.Open("data.txt")
	if openErr != nil {
		panic(openErr)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		commands = append(commands, strings.TrimSuffix(scanner.Text(), "\n"))
	}
	return commands
}
