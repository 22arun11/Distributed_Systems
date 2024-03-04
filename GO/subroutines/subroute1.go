package main

import (
    "fmt"
    "net/http"
)

// HandleIndex is a subroutine that handles HTTP requests to the index route ("/")
func HandleIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the homepage!")
}

func main() {
    http.HandleFunc("/", HandleIndex) // Register the HandleIndex subroutine to be called on the "/" route

    fmt.Println("Starting server on port 8080")
    http.ListenAndServe(":8080", nil) // Start the server
}