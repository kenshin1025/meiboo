package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	fmt.Println("Starting server at 'http://localhost:8080'")

	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}
