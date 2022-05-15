package main

import (
	"Task_Home24/parser"
	"fmt"
	"net/http"
)

// main is the entry point of the program.
func main() {

	url := "https://www.w3schools.com/html/html_forms.asp"
	fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)

	// handle the error if there is one
	if err != nil {
		panic(err)
	}

	// doing this now so it won't be forgotten
	defer resp.Body.Close()

	err = parser.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(parser.GetLinks())
	fmt.Println(parser.GetHeadings())
	fmt.Println(parser.GetForm())

}
