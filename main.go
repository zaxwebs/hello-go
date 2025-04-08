package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	fmt.Fprintln(w, "Hello from Go!")
}

func handlerX(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/x" {
		notFoundHandler(w, r)
		return
	}
	fmt.Fprintln(w, "Hello from Go - X!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		notFoundHandler(w, r)
		return
	}
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current time: %s", currentTime)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Custom 404 - Page not found.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/x", handlerX)
	http.HandleFunc("/time", timeHandler)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
