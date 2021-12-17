package challenges

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func loadGame() (boards []board, numbers []int) {
	content := readInputOfDay(4)
	inputs := strings.Split(content, "\n\n")
	numberInputs := strings.Split(inputs[0], ",")
	numbers = make([]int, len(numberInputs))
	for index, value := range numberInputs {
		numbers[index] = str2int(value)
	}
	boardsCount := len(inputs) - 1
	boardSize := int(math.Sqrt(float64(len(strings.Fields(inputs[1])))))
	boards = make([]board, boardsCount)
	// fill boards with the given input
	for i := 0; i < boardsCount; i++ {
		rows := strings.Split(inputs[i+1], "\n")
		boards[i].size = boardSize
		boards[i].content = make([][]boardCell, boardSize)
		boards[i].locations = make(map[int]location)
		boards[i].stats.rowsChecked = make([]int, boardSize)
		boards[i].stats.colsChecked = make([]int, boardSize)
		for x := 0; x < boardSize; x++ {
			data := strings.Fields(rows[x])
			boards[i].content[x] = make([]boardCell, boardSize)
			for y := 0; y < boardSize; y++ {
				boards[i].content[x][y].amount = str2int(data[y])
				boards[i].locations[boards[i].content[x][y].amount] = location{x, y}
			}
		}
	}
	return
}

func Day_4_first() {
	boards, numbers := loadGame()
	var wonBoard *board

outer: // label for the outer loop
	for _, number := range numbers {
		for _, b := range boards {
			if b.check(number) {
				wonBoard = &b
				break outer // break the outer loop
			}
		}
	}

	fmt.Println(wonBoard)
	os.Exit(0)
}

func Day_4_second() {
	boards, numbers := loadGame()
	wonBoards := make(map[int]bool)

outer: // label for the outer loop
	for _, number := range numbers {
		for index, b := range boards {
			if b.check(number) {
				wonBoards[index] = true
				if len(wonBoards) == len(boards) {
					fmt.Println(b)
					break outer // break the outer loop
				}
			}
		}
	}

	os.Exit(0)
}
