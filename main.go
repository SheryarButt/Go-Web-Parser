package main

import (
	"Task_Home24/parser"
	"strings"
)

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page1">A link to another page1</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	err := parser.Parse(r)
	if err != nil {
		panic(err)
	}
}
