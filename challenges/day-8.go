package challenges

import (
	"fmt"
	"strings"
)

type segmentIO struct {
	inputs  []string
	outputs []string
}

func splitIO(s string) (sio segmentIO) {
	splitted := strings.Split(s, " | ")
	sio.inputs = strings.Fields(splitted[0])
	sio.outputs = strings.Fields(splitted[1])
	return
}

func parseSegmentIO(segmentIO string) (s []segmentIO) {
	segmentIOList := strings.Split(segmentIO, "\n")
	segmentIOList = segmentIOList[:len(segmentIOList)-1]
	for _, v := range segmentIOList {
		s = append(s, splitIO(v))
	}
	return
}

func Day_8_first() {
	content := readInputOfDay(8)
	uniqueOutputsCount := 0
	for _, v := range parseSegmentIO(content) {
		for _, output := range v.outputs {
			if len(output) == 2 || len(output) == 3 || len(output) == 4 || len(output) == 7 {
				uniqueOutputsCount++
			}
		}
	}
	fmt.Println(uniqueOutputsCount)
}
