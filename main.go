package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AddRequest struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}

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

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not allowed. Please use POST.")
		return
	}

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	result := AddResponse{
		Result: req.Number1 + req.Number2,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Custom 404 - Page not found.")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/x", handlerX)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/add", addHandler)

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
