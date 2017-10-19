package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("using port %s", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
