package main

import (
    "net/http"
    "github.com/alabiemmanuel177/url-shortener/handlers"
)

func main() {
    http.HandleFunc("/", handlers.HelloHandler)
    http.ListenAndServe(":8080", nil)
}