package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSON: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func ResponseError(w http.ResponseWriter, statusCode int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}

	if statusCode > 499 {
		log.Printf("Responding with 5XX error: %s\n", msg)
	}

	type response struct {
		Error string `json:"error"`
	}

	Response(w, statusCode, response{
		Error: msg,
	})
}
