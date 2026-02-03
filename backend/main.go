package main

import (
	"encoding/json"
	"errors"
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

// Error definitions

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInvalidFormat   = errors.New("invalid format")
	ErrInternalFailure = errors.New("internal failure")
)

func processAsciiRequest(r *http.Request) (int, string, error) {
	var req asciiRequest

	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println("[ Status: ", http.StatusBadRequest, "] ERROR: invalid JSON body:", err)
			return http.StatusBadRequest, "", ErrBadRequest
		}
	} else {
		req.Text = r.URL.Query().Get("text")
		req.Format = r.URL.Query().Get("format")
	}

	if req.Text == "" {
		log.Println("[ Status: ", http.StatusBadRequest, "] ERROR: no text provided")
		return http.StatusBadRequest, "", ErrBadRequest
	}

	if req.Format == "" {
		req.Format = "standard.txt"
	}

	printFormat := process.Wrapper(req.Format)
	if len(printFormat) == 0 {
		log.Println("[ Status: ", http.StatusInternalServerError, "] ERROR: failed to load font format")
		return http.StatusInternalServerError, "", ErrInvalidFormat
	}

	result := print.AsciiArt(req.Text, printFormat)
	return http.StatusOK, result, nil
}

func asciiWeb(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/" || r.Method != http.MethodGet:
		http.ServeFile(w, r, "../frontend/404.html")
		log.Println("[ Status: ", http.StatusNotFound, "] ERROR: invalid request path or method")
		return
	}
	// the first request to load the web page
	http.ServeFile(w, r, "../frontend/index.html")

	// Example: you decide to call ASCII logic from here later
	status, _, err := processAsciiRequest(r)
	if err != nil {
		if status == http.StatusHTTPVersionNotSupported {
			http.ServeFile(w, r, "../frontend/error.html")
			log.Println("[ Status: ", http.StatusHTTPVersionNotSupported, "] ERROR: HTTP version not supported")
			return
		}

		http.ServeFile(w, r, "../frontend/error.html")
		return
	}

	http.ServeFile(w, r, "../frontend/index.html")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	status, result, err := processAsciiRequest(r)
	if err != nil {
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": result,
	})
	log.Println("[ Status: ", http.StatusOK, "] INFO: ASCII art generated successfully")
}

func main() {
	http.HandleFunc("/", asciiWeb)
	http.HandleFunc("/ascii-art", testHandler)
	fmt.Println("INFO: server running on port 8000, check http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
