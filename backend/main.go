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

type RequestData struct {
	Text string `json:"text"`
	Font string `json:"font"`
}



/*

func errorCheck(err error) {
	if err != nil {
		log.Fatal("ERROR:	", err)
	}
}
*/

func stringCheck(s string) {
	if s == "" {
		log.Fatal("ERROR:		input string is empty")
		return
	}
	log.Println("INFO:		input string is valid")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    // Allow any origin (for dev purposes)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

    // Handle preflight request
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    var req RequestData
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Text == "" {
        http.Error(w, "Text cannot be empty", http.StatusBadRequest)
        return
    }

    fontFile := req.Font
    if fontFile == "" {
        fontFile = "standard.txt"
    }

    printFormat := process.Wrapper(fontFile)
    if printFormat == nil || len(printFormat) == 0 {
        http.Error(w, "Failed to load font", http.StatusInternalServerError)
        return
    }

    result := print.AsciiArt(req.Text, printFormat)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": result})
}


func main() {
	http.HandleFunc("/", testHandler)
	fmt.Println("INFO:		server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
