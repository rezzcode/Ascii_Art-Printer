package main

import (
	"os"
	"strings"
	"testing"

	"ascii_art/Lib/print"
	"ascii_art/Lib/process"
)

func TestAsciiArt(t *testing.T) {
	testCases := []struct {
		inputfile          string
		expectedoutputfile string
	}{
		{"Tests/input1.txt", "Tests/output1.txt"},
		{"Tests/input2.txt", "Tests/output2.txt"},
		{"Tests/input3.txt", "Tests/output3.txt"},
		{"Tests/input4.txt", "Tests/output4.txt"},
		{"Tests/input5.txt", "Tests/output5.txt"},
		{"Tests/input7.txt", "Tests/output7.txt"},
		{"Tests/input8.txt", "Tests/output8.txt"},
		{"Tests/input9.txt", "Tests/output9.txt"},
		{"Tests/input10.txt", "Tests/output10.txt"},
		{"Tests/input11.txt", "Tests/output11.txt"},
		{"Tests/input12.txt", "Tests/output12.txt"},
		{"Tests/input13.txt", "Tests/output13.txt"},
		{"Tests/input14.txt", "Tests/output14.txt"},
		{"Tests/input15.txt", "Tests/output15.txt"},
	}

	for _, tc := range testCases {
		input, err := os.ReadFile(tc.inputfile)
		if err != nil {
			t.Fatalf("Failed to read input file %s: %v", tc.inputfile, err)
		}

		expectedOutput, err := os.ReadFile(tc.expectedoutputfile)
		if err != nil {
			t.Fatalf("Failed to read expected output file %s: %v", tc.expectedoutputfile, err)
		}

		asciitemplates := process.Wrapper("standard.txt")

		// call function
		output := print.AsciiArt(string(input), asciitemplates)

		// compare actual output to expected output
		if strings.TrimSpace(output) != strings.TrimSpace(string(expectedOutput)) {
			t.Errorf(
				"Output mismatch for %s\nExpected:\n%s\nGot:\n%s",
				tc.inputfile,
				expectedOutput,
				output,
			)
		}
	}
}
