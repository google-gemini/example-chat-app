package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

func parseRequestJSON(r *http.Request, target any) error {
	contentType := r.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expecting application/json Content-Type. Got %s", mediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	return dec.Decode(target)
}

func renderResponseJSON(w http.ResponseWriter, res any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
