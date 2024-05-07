package main

import (
	"io"
	"log"
	"fmt"
	"net/http"
)

// Greet receives a writer & string to print
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Handler function to be used as a http request handler
// simply calls the Greet function, uses to print a string
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	// line to check that the server is running
	fmt.Print("running on port 5001 \n")

	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}