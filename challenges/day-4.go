package challenges

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// board's cell data
type boardCell struct {
	amount  int
	checked bool
}

// board's statistics
type boardStats struct {
	rowsChecked []int
	colsChecked []int
	won         bool
	wonOnNumber int
	score       int
	finalScore  int
}

// location of an element in the board
type location struct {
	x int
	y int
}

// board is a representation of a bingo board
type board struct {
	locations map[int]location
	content   [][]boardCell
	stats     boardStats
	size      int
}

// report the bingo board's stats
func (b board) String() string {
	return fmt.Sprintf("won on: %d\nscore: %d\nfinal score: %d\n", b.stats.wonOnNumber, b.stats.score, b.stats.finalScore)
}

// check if the bingo board is won
func (b *board) checkBoard() {
	for _, row := range b.stats.rowsChecked {
		if row == b.size {
			b.stats.won = true
		}
	}
	for _, col := range b.stats.colsChecked {
		if col == b.size {
			b.stats.won = true
		}
	}
}

// calculate the score of the bingo board
func (b *board) calculateScore() {
	for _, row := range b.content {
		for _, data := range row {
			if !data.checked {
				b.stats.score += data.amount
			}
		}
	}
	b.stats.finalScore = b.stats.score * b.stats.wonOnNumber
}

// check for a number in the bingo board
func (b *board) check(number int) (gameOver bool) {
	if loc, found := b.locations[number]; found {
		b.content[loc.x][loc.y].checked = true
		b.stats.rowsChecked[loc.x]++
		b.stats.colsChecked[loc.y]++
		if b.checkBoard(); b.stats.won {
			b.stats.wonOnNumber = number
			b.calculateScore()
			gameOver = true
		}
	}
	return
}

func Day_4_first() {
	content := readInputOfDay(4)
	inputs := strings.Split(content, "\n\n")
	numberInputs := strings.Split(inputs[0], ",")
	numbers := make([]int, len(numberInputs))
	for index, value := range numberInputs {
		numbers[index] = str2int(value)
	}
	boardsCount := len(inputs) - 1
	boardSize := int(math.Sqrt(float64(len(strings.Fields(inputs[1])))))
	boards := make([]board, boardsCount)
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
	// we should change this part of the code to be able to solve the second part
	// of the challenge.
	// play the game
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
