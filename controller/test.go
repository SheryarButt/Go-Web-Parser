package controller

import (
	"Task_Home24/views"
	"encoding/json"
	"net/http"
)

// ping handles the request for the web page.
func helloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := views.Response{
				StatusCode: http.StatusOK,
				Body:       "Hello World, Please enter the 'url' in the query string over /getDetails or /getCounts endpoint",
			}
			err := json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
