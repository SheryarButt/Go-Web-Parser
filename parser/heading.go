package parser

import (
	"sync"

	"golang.org/x/net/html"
)

// Headings is a struct that represents a heading in the web page.
type Headings struct {
	Level string
	Text  string
}

// parseHeadings parses the heading node and adds it to the ParsedInformation struct.
func parseHeadings(n *html.Node, parsed *ParsedInformation, wg *sync.WaitGroup) {
	if string(n.Data[0]) == "h" {
		heading := Headings{
			Level: n.Data,
			Text:  getText(n),
		}
		parsed.Headings = append(parsed.Headings, heading)
	}
	wg.Done()
}

// GerHeadingCounts returns the number of each headings found in the web page to the ParsedInformation struct.
func GetHeadingCount(parsed *ParsedInformation) {

	for _, h := range parsed.Headings {
		switch h.Level {
		case "h1":
			parsed.Count.H1++
		case "h2":
			parsed.Count.H2++
		case "h3":
			parsed.Count.H3++
		case "h4":
			parsed.Count.H4++
		case "h5":
			parsed.Count.H5++
		case "h6":
			parsed.Count.H6++
		}
	}
}
