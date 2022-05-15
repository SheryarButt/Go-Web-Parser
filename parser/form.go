package parser

import "golang.org/x/net/html"

// hasForm is a global variable that stores whether the web page contains a form.
var hasForm bool

// parseForm parses the form node and sets the global variable hasForm to true.
func parseForm(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "form" {
		hasForm = true
	}
}

// getForm returns the form found in the web page.
func GetForm() bool {
	return hasForm
}
