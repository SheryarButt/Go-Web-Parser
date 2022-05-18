package parser

import (
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

// Link represents a link in a web page.
// Sample: <a href="...."> link text </a>
type Link struct {
	Href   string
	Text   string
	Type   string
	Status string
}

// parseLink parses the link node and adds it to the ParsedInformation struct.
func parseLink(n *html.Node, parsed *ParsedInformation, wg *sync.WaitGroup) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			link := Link{
				Href:   a.Val,
				Text:   getText(n),
				Type:   getType(a.Val, parsed),
				Status: getStatus(a.Val, parsed),
			}
			parsed.Links = append(parsed.Links, link)
		}
	}
	wg.Done()
}

// getType returns the type of the link.
func getType(s string, parsed *ParsedInformation) string {
	if len(s) > 0 {
		if s[0] == 'h' {
			return "External"
		}
		parsed.Count.InternalLinks++
	}
	return "Internal"
}

// getStatus returns the status of the link.
func getStatus(s string, parsed *ParsedInformation) string {
	if len(s) > 0 {
		if s[0] == 'h' {
			_, err := http.Get(s)
			if err == nil {
				parsed.Count.ExternalActiveLinks++
				return "Accessible"
			}
			parsed.Count.ExternalDeadLinks++
			return "Inaccessible"
		}
	}
	return "Internal" // Not performing accessibility check for internal links.
}
