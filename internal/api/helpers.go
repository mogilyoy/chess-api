package api

import (
	"encoding/json"
	"net/http"

	"github.com/oklog/ulid/v2"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func NewULID() ULID {
	return ulid.Make()
}
