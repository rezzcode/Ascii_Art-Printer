package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii_art/Lib/process"
)

func main() {
	http.HandleFunc("/", process.AsciiWeb)
	http.HandleFunc("/ascii-art", process.TestHandler)
	fmt.Println("INFO: server running on port 8000, check http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
