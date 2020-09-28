package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HelloMessage(w)
	})
	http.ListenAndServe(":80", nil)
}

// HelloMessage prints hello, world
func HelloMessage(w io.Writer) {
	fmt.Fprintf(w, "Hello, World")
}
