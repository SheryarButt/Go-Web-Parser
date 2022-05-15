package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parse parses the web page and returns error if the parsing fails.
func Parse(r io.Reader) (err error) {

	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	depthFirstSearch(doc)
	return nil
}

// depthFirstSearch traverses the web page and calls parseLink for each link node.
func depthFirstSearch(n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			parseLink(n)
		case "h1", "h2", "h3", "h4", "h5", "h6":
			parseHeadings(n)
		case "form":
			parseForm(n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		depthFirstSearch(c)
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
