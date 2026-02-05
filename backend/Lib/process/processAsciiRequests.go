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


// ================= LOGIC LAYER =================
// NO http.ResponseWriter
// NO *http.Request
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


func AsciiWeb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.ServeFile(w, r, "../frontend/404.html")
		log.Println("[404] invalid path or method")
		return
	}

	http.ServeFile(w, r, "../frontend/index.html")
}


func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "POST method required", http.StatusBadRequest)
		return
	}

	var req asciiRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Safety net → catches panics → 500
	defer func() {
		if rec := recover(); rec != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("[500] panic:", rec)
		}
	}()

	result, err := asciiLogic(req)
	if err != nil {
		switch err {
		case ErrBadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Status Bad Request",http.StatusBadRequest)
		case ErrInvalidFormat:
			http.Error(w, err.Error(), http.StatusNotFound)
			log.Println("Status not found",http.StatusNotFound)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			log.Println("Internal server error",http.StatusInternalServerError)

		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": result,
	})

	log.Println("[200] ASCII generated successfully")
}
