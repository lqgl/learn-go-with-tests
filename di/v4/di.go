package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	// Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":8000", http.HandlerFunc(MyGreetHandler))
}
