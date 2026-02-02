package main

import (
	"fmt"

	//"strconv"
	"encoding/json"
	"log"
	"net/http"

	"ascii_art/Lib/print"
	"ascii_art/Lib/process"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	/*
		input := os.Args

		fileName, data, err := check.Args(input)

		if !err {
			fmt.Println(data)
			return
		}
	*/
	// still needs to be tested with a [[[],[]],[[],[]]]
	data := "T"
	printFormat := process.Wrapper("standard.txt")

	// printFormat := process.Wrapper(fileName)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("INFO:		status code = ", http.StatusOK)

	// print to web console
	result := print.AsciiArt(data, printFormat)
	json.NewEncoder(w).Encode(map[string]string{"message": result})
}

func main() {
	http.HandleFunc("/", testHandler)
	fmt.Println("INFO:		server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
