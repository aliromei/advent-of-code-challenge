package challenges

import (
	"fmt"
	"os"
	"strings"
)

func Day_3_first() {
	content := readInputOfDay(3)
	inputs := strings.Fields(content)
	inputLength := len(inputs[0])
	var gamma, epsilon int
	result := make([]int, inputLength)
	for _, input := range inputs {
		for i := 0; i < inputLength; i++ {
			if input[i] == '1' {
				result[i]++
			} else {
				result[i]--
			}
		}
	}
	format := fmt.Sprintf("%%0%db\n", inputLength)
	gamma = strBitArr2Int(result)
	epsilon = (1 << uint(inputLength)) - gamma - 1
	fmt.Printf("ga: "+format, int(gamma))
	fmt.Printf("ep: "+format, int(epsilon))
	fmt.Printf("power consumption: %d\n", int(gamma*epsilon))
}

func Day_3_second() {
	content := readInputOfDay(3)
	inputs := strings.Fields(content)
	maxBit := uint(len(inputs[0]) - 1)
	convertedInputs := make([]int, len(inputs))

	for index, input := range inputs {
		convertedInputs[index] = binstr2int(input)
	}

	oxygenGeneratorRating := getFinder(convertedInputs, true, maxBit)
	co2ScrubberRating := getFinder(convertedInputs, false, maxBit)

	fmt.Println("Life support rating: ", oxygenGeneratorRating*co2ScrubberRating)
	os.Exit(0)
}

func getFinder(inputs []int, shouldFindGreaterOnes bool, maxBit uint) (finalResult int) {
	var result []int

	finder := func(nthBit uint) (found bool) {
		var ones, zeros []int
		for _, input := range inputs {
			if (input & (1 << nthBit)) != 0 {
				ones = append(ones, input)
			} else {
				zeros = append(zeros, input)
			}
		}
		condition := len(ones) < len(zeros)
		if shouldFindGreaterOnes {
			condition = len(ones) >= len(zeros)
		}
		if condition {
			result = ones
		} else {
			result = zeros
		}
		inputs = result
		if found = len(result) == 1; found {
			finalResult = result[0]
		}
		return
	}

	var found bool
	for nthBit := maxBit; !found; nthBit-- {
		found = finder(nthBit)
		if nthBit == 0 {
			break
		}
	}
	return
}
