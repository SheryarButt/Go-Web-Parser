package controller

import (
	"Task_Home24/parser"
	"Task_Home24/views"
	"encoding/json"
	"net/http"
	"net/url"
)

// processCounts handles the request from /getCounts endpoint.
func processCounts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Query().Get("url") != "" {
			URL, err := url.ParseRequestURI(r.URL.Query().Get("url"))
			if err != nil {
				data := views.Response{
					StatusCode: http.StatusBadRequest,
					Body:       "Bad Request Error",
				}
				err = json.NewEncoder(w).Encode(data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				resp, err := http.Get(URL.String())
				if err != nil {
					data := views.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       err.Error(),
					}
					err = json.NewEncoder(w).Encode(data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
				} else {
					parsed := new(parser.ParsedInformation)
					err := parser.Parse(resp.Body, parsed)
					if err != nil {
						data := views.Response{
							StatusCode: http.StatusInternalServerError,
							Body:       "Error while parsing HTML",
						}
						err = json.NewEncoder(w).Encode(data)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
					} else {
						data := views.SuccessResponse{
							HtmlVersion: parsed.Doctype.Version,
							Title:       parsed.Title,
							LoginForm:   parsed.LoginForm,
							Counts:      parsed.Count,
						}
						err = json.NewEncoder(w).Encode(data)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
					}
				}
			}
		} else if r.Method == http.MethodGet {
			data := views.Response{
				StatusCode: http.StatusOK,
				Body:       "Server is running, please use GET method with 'url' parameter",
			}
			err := json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// processDetails handles the request from /getDetails endpoint.
func processDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Query().Get("url") != "" {
			URL, err := url.ParseRequestURI(r.URL.Query().Get("url"))
			if err != nil {
				data := views.Response{
					StatusCode: http.StatusBadRequest,
					Body:       "Bad Request Error",
				}
				err = json.NewEncoder(w).Encode(data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			} else {
				resp, err := http.Get(URL.String())
				if err != nil {
					data := views.Response{
						StatusCode: http.StatusInternalServerError,
						Body:       err.Error(),
					}
					err = json.NewEncoder(w).Encode(data)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
				} else {
					parsed := new(parser.ParsedInformation)
					err := parser.Parse(resp.Body, parsed)
					if err != nil {
						data := views.Response{
							StatusCode: http.StatusInternalServerError,
							Body:       "Error while parsing HTML",
						}
						err = json.NewEncoder(w).Encode(data)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
					} else {
						data := views.Response{
							StatusCode: http.StatusOK,
							Body:       parsed,
						}
						err = json.NewEncoder(w).Encode(data)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
					}
				}
			}
		} else if r.Method == http.MethodGet {
			data := views.Response{
				StatusCode: http.StatusOK,
				Body:       "Server is running, please use GET method with 'url' parameter",
			}
			err := json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
