package parser

import (
	"golang.org/x/net/html"
)

// parseTitle parses the title node and updates the information in ParsedInformation struct.
func parseTitle(n *html.Node, title *string) {
	if n.Type == html.ElementNode && n.Data == "title" {
		*title = getText(n)
	}
}
