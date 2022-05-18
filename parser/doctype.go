package parser

import (
	"sync"

	"golang.org/x/net/html"
)

// Doctype is a struct that represents a doctype in the web page.
type Doctype struct {
	Value   string
	Version string
}

// receiver function to return the html version.
// func (d *Doctype) String() string {
// 	return d.Version
// }

// doctype is a global variable that stores the doctype.
// var doctype Doctype

// doctypes is a map that stores the doctype and its version.
var doctypes = make(map[string]string)

// init initializes the doctypes map.
func init() {
	doctypes["-//W3C//DTD HTML 4.01//EN"] = "HTML 4.01 Strict"
	doctypes["-//W3C//DTD HTML 4.01 Transitional//EN"] = "HTML 4.01 Transitional"
	doctypes["-//W3C//DTD HTML 4.01 Frameset//EN"] = "HTML 4.01 Frameset"
	doctypes["-//W3C//DTD XHTML 1.0 Strict//EN"] = "XHTML 1.0 Strict"
	doctypes["-//W3C//DTD XHTML 1.0 Transitional//EN"] = "XHTML 1.0 Transitional"
	doctypes["-//W3C//DTD XHTML 1.0 Frameset//EN"] = "XHTML 1.0 Frameset"
	doctypes["-//W3C//DTD XHTML 1.1//EN"] = "XHTML 1.1"
	doctypes["html"] = "HTML 5"
}

// docTypeParser parses the doctype node.
func docTypeParser(n *html.Node, parsed *ParsedInformation, wg *sync.WaitGroup) {
	var Value string
	if n.Attr != nil {
		var p string
		for _, a := range n.Attr {
			switch a.Key {
			case "public":
				p = a.Val
			}
		}
		if p != "" {
			Value += p
		}
	}
	if Value == "" {
		Value = n.Data
	}

	parsed.Doctype = Doctype{
		Value:   Value,
		Version: getDocType(Value),
	}
	wg.Done()
}

// getDocType returns the version of the doctype.
func getDocType(s string) string {

	val, ok := doctypes[s]
	if ok {
		return val
	}
	return "Unknown"
}

// GetDoctype returns the doctype of the web page.
// func GetDoctype() Doctype {
// 	return doctype
// }

// ifUnknown returns unknown doctype if it is nil.
func ifUnknown() Doctype {
	return Doctype{
		Value:   "Unknown",
		Version: "Unknown",
	}
}
