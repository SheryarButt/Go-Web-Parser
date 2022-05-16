package parser

import (
	"golang.org/x/net/html"
)

// title is a global variable that stores the title of the web page.
var title string

// parseTitle parses the title node and sets the global variable title to the title of the web page.
func parseTitle(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "title" {
		title = getText(n)
	}
}

// GetTitle returns the title of the web page.
func GetTitle() string {
	return title
}
