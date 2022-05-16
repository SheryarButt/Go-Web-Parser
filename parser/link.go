package parser

import (
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

// Link respresents a link in a web page.
// Sample: <a href="...."> link text </a>
type Link struct {
	Href   string
	Text   string
	Type   bool // true if it is an External link, false if it is an Internal link.
	Status bool // true if the link is valid, false if it is not valid.
}

// counters for the links found in the web page.
var internalLinks, externalActiveLinks, externalDeadLinks uint

// links is a global variable that stores all the links found in the web page.
var links []Link

// parseLink parses the link node and adds it to the links slice.
func parseLink(n *html.Node, wg *sync.WaitGroup) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			link := Link{
				Href:   a.Val,
				Text:   getText(n),
				Type:   getType(a.Val),
				Status: getStatus(a.Val),
			}
			links = append(links, link)
		}
	}
	wg.Done()
}

// getType returns true if the link is an external link, false if it is an internal link.
func getType(s string) bool {
	if len(s) > 0 {
		if s[0] == 'h' {
			return true
		}
		internalLinks++
	}
	return false
}

// getStatus returns true if the link is valid, false if it is not valid.
func getStatus(s string) bool {
	if len(s) > 0 {
		if s[0] == 'h' {
			_, err := http.Get(s)
			if err == nil {
				externalActiveLinks++
				return true
			}
			externalDeadLinks++
		}
	}
	return false
}

// GetLinks returns the links found in the web page.
func GetLinks() []Link {
	return links
}

// GetLinkCount returns the number of links found in the web page.
func GetLinkCount() (uint, uint, uint) {
	return internalLinks, externalActiveLinks, externalDeadLinks
}
