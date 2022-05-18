package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	go mux.HandleFunc("/ping", ping())
	go mux.HandleFunc("/", process())

	return mux
}
