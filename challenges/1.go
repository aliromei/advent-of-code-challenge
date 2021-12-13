package challenges

import (
	"fmt"
	"os"
	"strings"
)

// runs the 1st challenge
func First() {
	// read the file
	data, err := os.ReadFile("inputs/1")
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
