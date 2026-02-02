package main

import (
	"fmt"
	"os"

	//"strconv"

	"ascii_art/Lib/process"
	"ascii_art/Lib/check"
	"ascii_art/Lib/print"
)

func main() {
	input := os.Args

	fileName, data, err := check.Args(input)

	if !err {
		fmt.Println(data)
		return
	}
	// still needs to be tested with a [[[],[]],[[],[]]]
	// printFormat := Lib.Wrapper("standard.txt")

	printFormat := process.Wrapper(fileName)

	print.AsciiArt(data, printFormat)
}
