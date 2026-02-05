package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"ascii_art/Lib/process"
)

func TestProcessAsciiRequest(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		query          string
		expectedStatus int
		expectError    bool
	}{
		{
			name:           "valid POST request",
			method:         http.MethodPost,
			body:           `{"text":"Hi","format":"standard.txt"}`,
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "invalid JSON body",
			method:         http.MethodPost,
			body:           `{"text":`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "missing text field",
			method:         http.MethodPost,
			body:           `{"format":"standard.txt"}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "GET fallback with query params",
			method:         http.MethodGet,
			query:          "?text=Hello&format=standard.txt",
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:           "invalid format file",
			method:         http.MethodPost,
			body:           `{"text":"Hello","format":"does-not-exist.txt"}`,
			expectedStatus: http.StatusInternalServerError,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request

			if tt.method == http.MethodPost {
				req = httptest.NewRequest(
					tt.method,
					"/ascii-art",
					bytes.NewBufferString(tt.body),
				)
			} else {
				req = httptest.NewRequest(
					tt.method,
					"/ascii-art"+tt.query,
					nil,
				)
			}

			status, result, err := process.AsciiRequest(req)

			if status != tt.expectedStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.expectedStatus,
					status,
				)
			}

			if tt.expectError && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tt.expectError && err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}

			if !tt.expectError && result == "" {
				t.Fatalf("expected ASCII result, got empty string")
			}
		})
	}
}
