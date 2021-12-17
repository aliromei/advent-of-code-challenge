package challenges

import (
	"strings"
)

func Day_2_first() {
	content := readInputOfDay(2)
	// result to store the result
	result := map[string]int{"x": 0, "y": 0}
	// tokenized inputs from the inputs file
	inputs := strings.Split(content, "\n")
	// loop through the inputs
	for _, input := range inputs {
		// tokenize the input
		values := strings.Split(input, " ")
		if len(values) != 2 {
			continue
		}
		amount := str2int(values[1])
		switch values[0] {
		case "forward":
			result["x"] += amount
		case "up":
			result["y"] -= amount
		case "down":
			result["y"] += amount
		}
	}
	// print the result
	println(result["x"] * result["y"])
}

func Day_2_second() {
	content := readInputOfDay(2)
	// result to store the result
	result := map[string]int{"x": 0, "y": 0, "aim": 0}
	// tokenized inputs from the inputs file
	inputs := strings.Split(content, "\n")
	// loop through the inputs
	for _, input := range inputs {
		// tokenize the input
		values := strings.Split(input, " ")
		if len(values) != 2 {
			continue
		}
		amount := str2int(values[1])
		switch values[0] {
		case "forward":
			result["x"] += amount
			result["y"] += amount * result["aim"]
		case "up":
			result["aim"] -= amount
		case "down":
			result["aim"] += amount
		}
	}
	// print the result
	println(result["x"] * result["y"])
}
