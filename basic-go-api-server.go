package main

/**
This is a basic Go API server that echos back the request body as response.

To run the server, use the following command:
	go run basic-go-api-server.go

To test the server, go to http://localhost:8080
*/

import (
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

// main function to start the server
func main() {
	// Register the echo handler for all routes
	http.HandleFunc("/", helloWorldHandler)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server error:", err)
	}
}
