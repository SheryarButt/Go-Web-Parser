package parser

import (
	"sync"

	"golang.org/x/net/html"
)

// hasForm is a global variable that stores whether the web page contains a form.
var hasForm bool

// parseForm parses the form node and sets the global variable hasForm to true.
func parseForm(n *html.Node, wg *sync.WaitGroup) {
	if n.Type == html.ElementNode && n.Data == "form" {
		findLogin(n)
	}
	wg.Done()
}

// findLogin finds password field in form to check if the form contains a login form.
func findLogin(n *html.Node) {
	if n.Attr != nil {
		for _, a := range n.Attr {
			if a.Key == "type" && a.Val == "hidden" { // Using hidden type to identify password field in the form.
				hasForm = true
				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLogin(c)
	}
}

// getForm returns the form found in the web page.
func GetForm() bool {
	return hasForm
}
