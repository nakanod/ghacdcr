package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	message := "Hello, world on " + os.Getenv("ENV")
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
