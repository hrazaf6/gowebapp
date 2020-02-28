package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check</h1>")
}

func check1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check1</h1>")
}

func check2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health Check2</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	http.HandleFunc("/health_check1", check1)
	http.HandleFunc("/health_check2", check2)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}