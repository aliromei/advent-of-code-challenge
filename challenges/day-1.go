package challenges

import (
	"fmt"
	"os"
	"strings"
)

// converts the given numeric string to its integer value
func str2int(s string) int {
	var i int
	if _, err := fmt.Sscanf(s, "%d", &i); err != nil {
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

// runs the 1st challenge
func First() {
	// read the file
	data, err := os.ReadFile("inputs/day-1")
	check(err)
	// tokenize the file
	inputs := strings.Fields(string(data))
	inputsLength, firstInput, counter := len(inputs), 0, 0
	if len(inputs) < 2 {
		fmt.Println("Insufficiant inputs")
		os.Exit(1)
	}
	_, err = fmt.Sscanf(inputs[0], "%d", &firstInput)
	check(err)
	// start the loop
	for a, b, i := 2000000, firstInput, 0; i < inputsLength-1; i++ {
		if a < b {
			counter++
		}
		a = b
		_, err = fmt.Sscanf(inputs[i+1], "%d", &b)
		check(err)
	}
	// print the result
	fmt.Println("increment count =", counter)
	os.Exit(0)
}

// runs the second challenge
func Second() {
	// read the file
	data, err := os.ReadFile("inputs/day-1")
	check(err)
	// tokenize the file
	inputs := strings.Fields(string(data))
	inputsLength, counter := len(inputs), 0
	if len(inputs) < 2 {
		fmt.Println("Insufficiant inputs")
		os.Exit(1)
	}
	// start the loop
	for a, b, i := sumOfStrings(inputs[0:3]), sumOfStrings(inputs[1:4]), 2; i < inputsLength-1; i++ {
		if a < b {
			counter++
		}
		if i+3 <= inputsLength {
			a, b = b, sumOfStrings(inputs[i:i+3])
		}
	}
	// print the result
	fmt.Println("increment count =", counter)
	os.Exit(0)
}
