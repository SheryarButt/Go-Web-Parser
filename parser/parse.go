package parser

import (
	"fmt"
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
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(c.Data)
		depthFirstSearch(c)
	}
}
