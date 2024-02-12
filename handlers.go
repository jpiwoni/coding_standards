package main

import (
	"fmt"
	"net/http"
)

func handlers() {
	http.HandleFunc("/", index)
	http.HandleFunc("/html", html)
	http.HandleFunc("/json", json)
	http.HandleFunc("/syllabus", syllabus)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/help", help)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World - GWS")
}

func html(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Hello World &mdash; GWS</h1>")
}

func json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"Hello World - GWS"}`)
}

// Questionable
func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete-stubbed")
}

// Questionable
func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create-stubbed")
}

// Questionable
func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update-stubbed")
}

// Questionable
func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `

	Step 1: How to start the Web Server 
	  go run .                Runs all go files and starts the web server
	  go run . gws help       Shows this help message in the command line
	
	Step 2:
	  Visit http://localhost:8080 after starting the server
	
	Step 3: Using the Web Server and its specific endpoints
	  /index                  Responds in plain text
	  /html                   Responds with HTML content
	  /json                   Responds with JSON content
	  /syllabus               Responds with the embedded syllabus in JSON format
	  /create                 Responds with "create-stubbed"
	  /update                 Responds with "update-stubbed"
	  /delete                 Responds with "delete-stubbed"
	  /help                   Shows this help message
	  
	  `)
}		
