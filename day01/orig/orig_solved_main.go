package orig

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
	commands := loadData()

	var (
		numTimesZero = 0
		dialVal      = 50
	)
	for _, ins := range commands {
		dialVal = next(dialVal, ins)
		if dialVal == 0 {
			numTimesZero += 1
		}
	}
	fmt.Printf("Num Times Hit Zero: %d, Final dial value: %d", numTimesZero, dialVal)
}

func splitInstruction(instruction string) (string, int) {
	amount, _ := strconv.Atoi(instruction[1:])
	return string(instruction[0]), amount
}

// max is 99, min is 0
func next(dialValue int, instruction string) int {
	direction, amount := splitInstruction(instruction)

	// rotates below 0
	if direction == "L" {
		dialValue = dialValue - amount
		for dialValue < 0 {
			dialValue = maxDialVal + dialValue // likely negative number at this point
		}
		// once negative turning has gone back to positive, fall through
	} else {
		// rotates past 99
		dialValue = dialValue + amount
		// only if greater than, since maxDialVal is also representative of 0
		for dialValue > maxDialVal {
			dialValue = dialValue - maxDialVal
		}
		// fall through
	}
	// maxDialValue is representative of 0
	if dialValue == maxDialVal {
		return 0
	}
	return dialValue // also returns zero correctly
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
