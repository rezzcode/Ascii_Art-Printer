package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ascii_art/Lib/print"
	"ascii_art/Lib/process"
)

type asciiRequest struct {
	Text   string `json:"text"`
	Format string `json:"format"`
}

func asciiWeb(w http.ResponseWriter, r *http.Request) {
	filePath := "../frontend/index.html"
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		filePath = "../frontend/404.html"
		http.ServeFile(w, r, filePath)
		http.NotFound(w, r)
		log.Println("ERROR: file not found, Status code:", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req asciiRequest

	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid JSON body"})
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			log.Println("ERROR: invalid JSON body:", err, "Status code:", http.StatusBadRequest)
			return
		}
	} else {
		// Allow simple GET fallback via query params: /test?text=Hi&format=standard.txt
		req.Text = r.URL.Query().Get("text")
		req.Format = r.URL.Query().Get("format")
	}

	if req.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "text is required"})
		http.Error(w, "Text is required", http.StatusBadRequest)
		log.Println("ERROR: text is required, Status code:", http.StatusBadRequest)
		return
	}

	if req.Format == "" {
		req.Format = "standard.txt"
	}

	log.Println("OPTIONS: received request", "text = ", req.Text, "format = ", req.Format)

	printFormat := process.Wrapper(req.Format)
	if len(printFormat) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR: Failed to load font format", http.StatusInternalServerError)
		return
	}

	// call existing library to produce ASCII
	result := print.AsciiArt(req.Text, printFormat)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": result})
	log.Println("INFO: successfully processed request, Status code:", http.StatusOK)
}

func main() {
	http.HandleFunc("/", asciiWeb)
	http.HandleFunc("/ascii-art", testHandler)
	fmt.Println("INFO: server running on port 8000, check http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
