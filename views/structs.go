package views

// Response is the response of a static request.
type Response struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}

// SuccessResponse is a struct that contains the data that is returned when the request is successful.
type SuccessResponse struct {
	HtmlVersion string      `json:"html_version"`
	Title       string      `json:"title"`
	Counts      interface{} `json:"counts"`
	LoginForm   bool        `json:"login_form"`
}
