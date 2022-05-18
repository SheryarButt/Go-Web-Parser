package main

import (
	"Task_Home24/controller"
	"fmt"
	"net/http"
)

// main is the entry point of the program.
func main() {

	mux := controller.Register()
	fmt.Println("Server is running on port 8080")
	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		panic(error)
	}
}
