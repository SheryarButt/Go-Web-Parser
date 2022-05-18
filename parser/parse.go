package parser

import (
	"io"
	"reflect"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type ParsedInformation struct {
	Doctype   Doctype
	Title     string
	LoginForm bool
	Links     []Link
	Headings  []Headings
	Count     Counts
}

type Counts struct {
	H1                  uint
	H2                  uint
	H3                  uint
	H4                  uint
	H5                  uint
	H6                  uint
	InternalLinks       uint
	ExternalActiveLinks uint
	ExternalDeadLinks   uint
}

// wg is a global variable that stores the wait group.
var wg sync.WaitGroup

// Parse parses the web page and returns error if the parsing fails.
func Parse(r io.Reader, parsed *ParsedInformation) (err error) {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	depthFirstSearch(doc, parsed)
	if reflect.DeepEqual(parsed.Doctype, Doctype{}) {
		parsed.Doctype = ifUnknown()
	}
	wg.Wait()
	GetHeadingCount(parsed)
	return nil
}

// depthFirstSearch traverses the web page and calls parseLink for each link node.
func depthFirstSearch(n *html.Node, parsed *ParsedInformation) {
	if n.Type == html.DoctypeNode {
		wg.Add(1)
		go docTypeParser(n, parsed, &wg)
	} else if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			wg.Add(1)
			go parseLink(n, parsed, &wg)
		case "h1", "h2", "h3", "h4", "h5", "h6":
			wg.Add(1)
			go parseHeadings(n, parsed, &wg)
		case "form":
			wg.Add(1)
			go parseForm(n, &parsed.LoginForm, &wg)
		case "title":
			go parseTitle(n, &parsed.Title)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		depthFirstSearch(c, parsed)
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
