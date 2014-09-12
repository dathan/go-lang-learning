package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my HTTP handler url : %s", r.URL.Path[1:])
}

func main() {
	fmt.Println("Staring server on http://localhost:8922")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8922", nil)
}
