package controller

import (
	"net/http"
)

// Register registers the endpoints to a mux and returns to the caller.
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	go mux.HandleFunc("/", helloWorld())
	go mux.HandleFunc("/getCounts", processCounts())
	go mux.HandleFunc("/getDetails", processDetails())

	return mux
}
