package challenges

import "fmt"

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
