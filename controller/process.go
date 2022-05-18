package controller

import (
	"Task_Home24/parser"
	"Task_Home24/views"
	"encoding/json"
	"net/http"
	"net/url"
)

// process handles the request for the web page.
func process() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Query().Get("URL") != "" {
			URL, err := url.ParseRequestURI(r.URL.Query().Get("URL"))
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

					err := parser.Parse(resp.Body)
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
						h1, h2, h3, h4, h5, h6 := parser.GetHeadingCount()
						internalLinks, externalActiveLinks, externalDeadLinks := parser.GetLinkCount()
						data := views.SuccessResponse{
							HtmlVersion:                parser.GetDoctype().Version,
							Title:                      parser.GetTitle(),
							H1Count:                    h1,
							H2Count:                    h2,
							H3Count:                    h3,
							H4Count:                    h4,
							H5Count:                    h5,
							H6Count:                    h6,
							InternalLinksCount:         internalLinks,
							ExternalActiveLinksCount:   externalActiveLinks,
							ExternalInActiveLinksCount: externalDeadLinks,
							LoginForm:                  parser.GetForm(),
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
				Body:       "Server is running, please use GET method with URL parameter",
			}
			err := json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
