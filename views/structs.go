package views

// Response is the response of a static request.
type Response struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}

// SuccessResponse is a struct that contains the data that is returned when the request is successful.
type SuccessResponse struct {
	HtmlVersion                string `json:"html_version"`
	Title                      string `json:"title"`
	H1Count                    uint   `json:"h1_count"`
	H2Count                    uint   `json:"h2_count"`
	H3Count                    uint   `json:"h3_count"`
	H4Count                    uint   `json:"h4_count"`
	H5Count                    uint   `json:"h5_count"`
	H6Count                    uint   `json:"h6_count"`
	InternalLinksCount         uint   `json:"internal_links_count"`
	ExternalActiveLinksCount   uint   `json:"external_active_links"`
	ExternalInActiveLinksCount uint   `json:"external_inactive_links"`
	LoginForm                  bool   `json:"login_form"`
}
