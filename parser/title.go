package parser

import (
	"golang.org/x/net/html"
)

// parseTitle parses the title node and sets the global variable title to the title of the web page.
func parseTitle(n *html.Node, title *string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		*title = getText(n)
	}
}
