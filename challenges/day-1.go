package challenges

import (
	"fmt"
	"os"
	"strings"
)

// runs the 1st challenge
func Day_1_first() {
	content := readInputOfDay(1)
	// tokenize the file
	inputs := strings.Fields(content)
	inputsLength, firstInput, counter := len(inputs), 0, 0
	if len(inputs) < 2 {
		fmt.Println("Insufficiant inputs")
		os.Exit(1)
	}
	if _, err := fmt.Sscanf(inputs[0], "%d", &firstInput); err != nil {
		panic(err)
	}
	// start the loop
	for a, b, i := 2000000, firstInput, 0; i < inputsLength-1; i++ {
		if a < b {
			counter++
		}
		a = b
		if _, err := fmt.Sscanf(inputs[i+1], "%d", &b); err != nil {
			panic(err)
		}
	}
	// print the result
	fmt.Println("increment count =", counter)
	os.Exit(0)
}

// runs the second challenge
func Day_1_second() {
	content := readInputOfDay(1)
	// tokenize the file
	inputs := strings.Fields(content)
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
