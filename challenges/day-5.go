package challenges

import (
	"fmt"
	"math"
	"strings"
)

type line struct {
	startPoint location
	endPoint   location
	dx         int
	dy         int
}

func readLine(coordinates string) (l line) {
	if _, err := fmt.Sscanf(coordinates, "%d,%d -> %d,%d", &l.startPoint.y, &l.startPoint.x, &l.endPoint.y, &l.endPoint.x); err != nil {
		panic(err)
	}
	l.dx = l.endPoint.x - l.startPoint.x
	l.dy = l.endPoint.y - l.startPoint.y
	return
}

func updateMaxSize(l line, maxSize *int) {
	if l.startPoint.x > *maxSize {
		*maxSize = l.startPoint.x
	}
	if l.startPoint.y > *maxSize {
		*maxSize = l.startPoint.y
	}
	if l.endPoint.x > *maxSize {
		*maxSize = l.endPoint.x
	}
	if l.endPoint.y > *maxSize {
		*maxSize = l.endPoint.y
	}
}

func Day_5_first() {
	content := readInputOfDay(5)
	straightLines := make([]line, 0)
	maxSize := 0
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		l := readLine(l)
		if l.dx == 0 || l.dy == 0 {
			straightLines = append(straightLines, l)
		}
		updateMaxSize(l, &maxSize)
	}
	grid := make([][]int, maxSize+1)
	for i := range grid {
		grid[i] = make([]int, maxSize+1)
	}
	for _, l := range straightLines {
		if l.dx == 0 {
			absDY := int(math.Abs(float64(l.dy)))
			for i := 0; i <= absDY; i++ {
				grid[l.startPoint.x][l.startPoint.y+i*(l.dy/absDY)] += 1
			}
		} else if l.dy == 0 {
			absDX := int(math.Abs(float64(l.dx)))
			for i := 0; i <= int(absDX); i++ {
				grid[l.startPoint.x+i*(l.dx/absDX)][l.startPoint.y] += 1
			}
		}
	}
	overlappingPoints := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] > 1 {
				overlappingPoints++
			}
		}
	}
	fmt.Println(overlappingPoints)
}

func Day_5_second() {
	content := readInputOfDay(5)
	lines := make([]line, 0)
	maxSize := 0
	for _, l := range strings.Split(content, "\n") {
		if l == "" {
			continue
		}
		l := readLine(l)
		if l.dx == 0 || l.dy == 0 || math.Abs(float64(l.dx)) == math.Abs(float64(l.dy)) {
			lines = append(lines, l)
		}
		updateMaxSize(l, &maxSize)
	}
	grid := make([][]int, maxSize+1)
	for i := range grid {
		grid[i] = make([]int, maxSize+1)
	}
	for _, l := range lines {
		if l.dx == 0 {
			absDY := int(math.Abs(float64(l.dy)))
			for i := 0; i <= absDY; i++ {
				grid[l.startPoint.x][l.startPoint.y+i*(l.dy/absDY)] += 1
			}
		} else if l.dy == 0 {
			absDX := int(math.Abs(float64(l.dx)))
			for i := 0; i <= int(absDX); i++ {
				grid[l.startPoint.x+i*(l.dx/absDX)][l.startPoint.y] += 1
			}
		} else if math.Abs(float64(l.dx)) == math.Abs(float64(l.dy)) {
			absD := int(math.Abs(float64(l.dx)))
			for i := 0; i <= absD; i++ {
				grid[l.startPoint.x+i*(l.dx/absD)][l.startPoint.y+i*(l.dy/absD)] += 1
			}
		}
	}
	overlappingPoints := 0
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] > 1 {
				overlappingPoints++
			}
		}
	}
	fmt.Println(overlappingPoints)
}
