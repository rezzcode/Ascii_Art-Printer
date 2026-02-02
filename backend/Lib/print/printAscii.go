package print

import (
	"fmt"
	"strings"
)

func indexToPrint(r rune) int {
	return int(r - 32)
}

func AsciiArt(data string, dataList [][][]rune) {
	data = strings.ReplaceAll(data, "\\n", "\n")
	lines := strings.Split(data, "\n")
	height := len(dataList[0])

	for _, line := range lines {
		rows := make([][]rune, height)
		if line == "" {
			fmt.Println()
			continue
		}

		for _, ch := range line {
			idx := indexToPrint(ch)
			visual := dataList[idx]

			for x := 0; x < height; x++ {
				rows[x] = append(rows[x], visual[x]...)
			}
		}

		for _, row := range rows {
			fmt.Println(string(row))
		}
		/**
		if endLine < len(lines)-2 {
			fmt.Println()
		}
		**/
	}
}
