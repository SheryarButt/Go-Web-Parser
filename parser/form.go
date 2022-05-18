package parser

import (
	"sync"

	"golang.org/x/net/html"
)

// parseForm parses the form node and updates the information in ParsedInformation struct.
func parseForm(n *html.Node, hasForm *bool, wg *sync.WaitGroup) {
	if n.Type == html.ElementNode && n.Data == "form" {
		findLogin(n, hasForm)
	}
	wg.Done()
}

// findLogin finds password field in form to check if the form contains a login form.
func findLogin(n *html.Node, hasForm *bool) {
	if n.Attr != nil {
		for _, a := range n.Attr {
			if a.Key == "type" && a.Val == "hidden" { // Using hidden type to identify password field in the form.
				*hasForm = true
				return
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLogin(c, hasForm)
	}
}
