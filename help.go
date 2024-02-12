package main

import "fmt"

func helpCL() {
	fmt.Println(`

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
