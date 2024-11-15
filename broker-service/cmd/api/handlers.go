package main

import (
	"net/http"
)


// http.ResponseWriter is an interface type, and in Go, interfaces are inherently <reference> types
// This means that even though it's not explicitly passed as a pointer, 
// you can still modify its underlying value (the response) within the handler function.
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the broker",
	}

  _ = app.writeJSON(w, payload, http.StatusOK)
}
