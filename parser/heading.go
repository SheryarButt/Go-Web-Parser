package parser

import (
	"sync"

	"golang.org/x/net/html"
)

// Headings is a struct that represents a heading in the web page.
type Headings struct {
	H    string
	Text string
}

// headings is a global variable that stores all the headings found in the web page.
var headings []Headings

// Global variables for counting each headings.
var h1, h2, h3, h4, h5, h6 uint

// parseHeadings parses the heading node and adds it to the headings slice.
func parseHeadings(n *html.Node, wg *sync.WaitGroup) {
	if string(n.Data[0]) == "h" {
		heading := Headings{
			H:    n.Data,
			Text: getText(n),
		}
		headings = append(headings, heading)
	}
	wg.Done()
}

// GetHeadings returns the headings found in the web page.
func GetHeadings() []Headings {
	return headings
}

// GerHeadingCounts returns the number of each headings found in the web page.
func GetHeadingCount() (uint, uint, uint, uint, uint, uint) {

	for _, h := range headings {
		switch h.H {
		case "h1":
			h1++
		case "h2":
			h2++
		case "h3":
			h3++
		case "h4":
			h4++
		case "h5":
			h5++
		case "h6":
			h6++
		}
	}
	return h1, h2, h3, h4, h5, h6
}
