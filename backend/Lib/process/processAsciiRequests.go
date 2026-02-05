package process

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"ascii_art/Lib/print"
)

type asciiRequest struct {
	Text   string `json:"text"`
	Format string `json:"format"`
}

// Errors (used by handler to map status codes)
var (
	ErrBadRequest      = errors.New("bad request")
	ErrInvalidFormat   = errors.New("invalid format")
	ErrInternalFailure = errors.New("internal failure")
)

func asciiLogic(req asciiRequest) (string, error) {
	if req.Text == "" {
		return "", ErrBadRequest
	}

	if req.Format == "" {
		req.Format = "standard.txt"
	}

	font := Wrapper(req.Format)
	if len(font) == 0 {
		return "", ErrInvalidFormat // 404
	}

	result := print.AsciiArt(req.Text, font)
	if result == "" {
		return "", ErrInternalFailure // 500
	}

	return result, nil
}

func AsciiRequest(r *http.Request) (int, string, error) {
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

	printFormat := Wrapper(req.Format)
	if len(printFormat) == 0 {
		log.Println("[ Status: ", http.StatusInternalServerError, "] ERROR: failed to load font format")
		return http.StatusInternalServerError, "", ErrInvalidFormat
	}

	result := print.AsciiArt(req.Text, printFormat)
	if result == "" {
		return http.StatusInternalServerError, "", ErrInternalFailure
	}
	return http.StatusOK, result, nil
}

func AsciiWeb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.ServeFile(w, r, "../frontend/404.html")
		log.Println("[ Status: ", http.StatusNotFound, "] ERROR: invalid request path or method")
		return
	}

	http.ServeFile(w, r, "../frontend/index.html")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "POST method required", http.StatusBadRequest)
		log.Println("[ Status: ", http.StatusBadRequest, "] ERROR: POST method required")
		return
	}

	var req asciiRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		log.Println("[ Status: ", http.StatusBadRequest, "] ERROR: invalid JSON body:", err)
		return
	}

	// Safety net → catches panics → 500
	defer func() {
		if rec := recover(); rec != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("[ Status: ", http.StatusInternalServerError, "] ERROR: internal server error:", rec)
		}
	}()

	result, err := asciiLogic(req)
	if err != nil {
		switch err {
		case ErrBadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("[ Status: ", http.StatusBadRequest, "] ERROR: bad request")
		case ErrInvalidFormat:
			http.Error(w, err.Error(), http.StatusNotFound)
			log.Println("[ Status: ", http.StatusNotFound, "] ERROR: invalid format")

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("[ Status: ", http.StatusInternalServerError, "] ERROR: internal server error:", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": result,
	})

	log.Println("[ Status: ", http.StatusOK, "] INFO: ASCII art generated successfully")
}
