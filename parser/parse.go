package parser

import (
	"io"

	"golang.org/x/net/html"
)

func Parse(r io.Reader) (err error) {

	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	depthFirstSearch(doc)
	return nil
}

func depthFirstSearch(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		parseLink(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		depthFirstSearch(c)
	}
}
