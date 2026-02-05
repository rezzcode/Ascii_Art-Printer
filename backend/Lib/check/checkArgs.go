package check

import (
	"os"
	"strings"
)

func Args(input []string) (string, string, bool) {
	inputLength := len(input)

	if inputLength < 2 {
		return "", "Did you pas a word like: go run . \"some-text?\"", false
	} else if inputLength > 3 {
		return "", "That's too much for me to handle", false
	}

	data := input[1]
	fileName := ""

	if len(data) < 1 {
		return "", "I can't work with an empty string", false
	}
	if inputLength == 2 {
		return "standard.txt", data, true
	}

	if inputLength == 3 {
		fileName = input[2]
		if len(fileName) < 1 {
			return "", "I can't work with an empty string", false
		}
		yString, err := FileEdgeCase(fileName)
		if !err {
			return "", yString, err
		}
		return fileName, data, true

	}

	return "", "Something realy bad off the check, happened", false
}

func FileEdgeCase(file string) (string, bool) {
	if !strings.HasSuffix(file, ".txt") {
		return "this is not the correct text file i want: should end with .txt", false
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return "I did not find that file, is it available?", false
	}

	if len(data) == 0 {
		return "This file is empty bruv!", false
	}

	return "exactly what i wanted", true
}
