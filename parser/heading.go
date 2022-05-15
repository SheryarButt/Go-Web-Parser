package parser

import (
	"golang.org/x/net/html"
)

// Headings is a struct that represents a heading in the web page.
type Headings struct {
	H    string
	Text string
}

// headings is a global variable that stores all the headings found in the web page.
var headings []Headings

// parseHeadings parses the heading node and adds it to the headings slice.
func parseHeadings(n *html.Node) {
	if string(n.Data[0]) == "h" {
		heading := Headings{
			H:    n.Data,
			Text: getText(n),
		}
		headings = append(headings, heading)
	}
}

// GetHeadings returns the headings found in the web page.
func GetHeadings() []Headings {
	return headings
}
