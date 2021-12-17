package challenges

import (
	"fmt"
	"os"
)

// reads the given file and returns the content
func readInputOfDay(day int) string {
	file, err := os.ReadFile(fmt.Sprintf("inputs/day-%d.txt", day))
	check(err)
	return string(file)
}

// on error print the error and exit
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// converts the given numeric string to its integer value
func str2int(s string) int {
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
		panic(err)
	}
	return i
}

// converts the given binary string to its integer value
func binstr2int(s string) int {
	var i int
	if _, err := fmt.Sscanf(s, "%b", &i); err != nil {
		panic(err)
	}
	return i
}

// gets an
func sumOfStrings(array []string) int {
	var sum int
	for _, value := range array {
		sum += str2int(value)
	}
	return sum
}

func strBitArr2Int(input []int) (result int) {
	inputLength := len(input)
	for i := 0; i < inputLength; i++ {
		if input[i] > 0 {
			result += 1 << uint(inputLength-1-i)
		}
	}
	return
}
