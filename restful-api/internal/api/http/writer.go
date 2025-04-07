package http

import (
	"fmt"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, resp []byte) {
    w.WriteHeader(status)
    
    if _, err := w.Write(resp); err != nil {
        WriteError(w, fmt.Errorf("failed to write response: %w", err))
    }
}
