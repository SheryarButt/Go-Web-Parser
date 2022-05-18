package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	go mux.HandleFunc("/", helloWorld())
	go mux.HandleFunc("/getCounts", processCounts())
	go mux.HandleFunc("/getDetails", processDetails())

	return mux
}
