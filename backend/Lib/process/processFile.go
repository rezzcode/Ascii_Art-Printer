package process

import (
	"bufio"
	"os"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	results := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results, nil
}

func ProcessResults(data []string) [][][]rune {
	const height = 8
	startInd := 1
	results := [][][]rune{}

	for startInd+height <= len(data) {
		visual := make([][]rune, height)

		for x := 0; x < height; x++ {
			visual[x] = []rune(data[startInd+x])
		}
		results = append(results, visual)
		startInd += height + 1
	}
	return results
}

func Wrapper(filename string) [][][]rune {
	results, err := readFile(filename)
	if err != nil {
		return nil
	}
	return ProcessResults(results)
}
