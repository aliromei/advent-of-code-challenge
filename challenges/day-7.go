package challenges

import (
	"fmt"
	"math"
	"strings"
)

func parseCsv(input string) (parsedInput []int) {
	for _, v := range strings.Split(input, ",") {
		parsedInput = append(parsedInput, str2int(v))
	}
	return
}

func maxInArray(arr []int) int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func minInArray(arr []int) int {
	min := 0
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

func calculateNeededFuel(fuelCalculator func(int) int) (min int) {
	content := readInputOfDay(7)
	fuelList := parseCsv(content)
	min = 1000000000
	for i := minInArray(fuelList); i <= maxInArray(fuelList); i++ {
		var fSum float64
		for _, v := range fuelList {
			fSum += float64(fuelCalculator(int(math.Abs(float64(v - i)))))
		}
		if int(fSum) < min {
			min = int(fSum)
		}
		fmt.Printf("%-10d %d\n", i, int(fSum))
	}
	return min
}

func Day_7_first() {
	calculator := func(distance int) (neededFuel int) {
		return distance
	}
	fmt.Println(calculateNeededFuel(calculator))
}

func Day_7_second() {
	calculator := func(distance int) (neededFuel int) {
		return distance * (distance + 1) / 2
	}
	fmt.Println(calculateNeededFuel(calculator))
}
