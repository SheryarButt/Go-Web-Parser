package parser

import (
	"strings"

	"golang.org/x/net/html"
)

// Link respresents a link in a web page.
// Sample: <a href="...."> link text </a>
type Link struct {
	Href string
	Text string
}

// links is a global variable that stores all the links found in the web page.
var links []Link

// parseLink parses the link node and adds it to the links slice.
func parseLink(n *html.Node) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			link := Link{
				Href: a.Val,
				Text: getText(n),
			}
			links = append(links, link)
		}
	}
}

// getText returns the text of the node.
func getText(n *html.Node) string {
	var b strings.Builder
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		b.WriteString(getText(c))
	}
	return strings.Join(strings.Fields(b.String()), " ")
}

// GetLinks returns the links found in the web page.
func GetLinks() []Link {
	return links
}
