package challenges

import (
	"fmt"
	"strings"
)

type line struct {
	startPoint location
	endPoint   location
	dx         int
	dy         int
}

func readLine(coordinates string) (l line) {
	if _, err := fmt.Sscanf(coordinates, "%d,%d -> %d,%d", &l.startPoint.x, &l.startPoint.y, &l.endPoint.x, &l.endPoint.y); err != nil {
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
		startPoint, endPoint := l.startPoint, l.endPoint
		if l.dx == 0 {
			if l.dy < 0 {
				startPoint, endPoint = l.endPoint, l.startPoint
			}
			for y := startPoint.y; y <= endPoint.y; y++ {
				grid[l.startPoint.x][y] += 1
			}
		} else if l.dy == 0 {
			if l.dx < 0 {
				startPoint, endPoint = l.endPoint, l.startPoint
			}
			for x := startPoint.x; x <= endPoint.x; x++ {
				grid[x][l.startPoint.y] += 1
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
