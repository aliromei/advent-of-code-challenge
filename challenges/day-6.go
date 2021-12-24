package challenges

import (
	"fmt"
	"strings"
)

func parseInput(input string) (parsedInput fishList) {
	for _, v := range strings.Split(input, ",") {
		parsedInput[str2int(v)]++
	}
	return
}

type fishList [9]int

func (fl *fishList) passADay() {
	flCopy := new(fishList)
	flCopy[8] = fl[0]
	flCopy[7] = fl[8]
	flCopy[6] = fl[7] + fl[0]
	flCopy[5] = fl[6]
	flCopy[4] = fl[5]
	flCopy[3] = fl[4]
	flCopy[2] = fl[3]
	flCopy[1] = fl[2]
	flCopy[0] = fl[1]
	for i := 0; i < 9; i++ {
		fl[i] = flCopy[i]
	}
}

func countFishesAfter(days int) int {
	content := readInputOfDay(6)
	fishes := fishList(parseInput(content))
	for i := 0; i < days; i++ {
		fishes.passADay()
		fmt.Println(fishes)
	}
	fishCount := 0
	for _, v := range fishes {
		fishCount += v
	}
	return fishCount
}

func Day_6_first() {
	fmt.Println(countFishesAfter(80))
}

func Day_6_second() {
	fmt.Println(countFishesAfter(256))
}
