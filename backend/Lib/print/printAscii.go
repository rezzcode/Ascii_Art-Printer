package print

import (
	"fmt"
	"log"
	"strings"
)

func indexToPrint(r rune) int {
	return int(r - 32)
}

func AsciiArt(data string, dataList [][][]rune) string {
	data = strings.ReplaceAll(data, "\\n", "\n")
	lines := strings.Split(data, "\n")

	if len(dataList) == 0 {
		log.Fatal("ERROR: dataList is empty. Cannot generate ASCII art.")
	}

	height := len(dataList[0])
	res := ""

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
			// fmt.Println(string(row)) uncomment this line to see output in terminal
			res += string(row) + "\n"
		}
		/**
		if endLine < len(lines)-2 {
			fmt.Println()
		}
		**/
	}
	return res
}
