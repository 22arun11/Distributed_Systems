package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

// FetchData is a subroutine that fetches data from a given URL
func FetchData(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Data fetched from the URL:")
    fmt.Println(string(body))
}

func main() {
    FetchData("https://www.google.com") // Replace with your URL
}