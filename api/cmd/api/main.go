package main

import (
	"fmt"
	"log"
	"net/http"
)

// customError defines a custom error type
type customError struct {
	code    int
	message string
}

// Error implements the error interface for customError
func (e *customError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.code, e.message)
}

// getRoot is the handler for the root path
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my website")
}

func main() {
	http.HandleFunc("/", getRoot)

	// Start the HTTP server on port 3000 and handle errors
	err := http.ListenAndServe(":3000", nil)
	fmt.Printf("Server started on port 3000")
	if err != nil {
		// Create a new custom error
		newError := &customError{code: 101, message: "Some error occurred"}

		// Log the custom error message using log.Fatalf
		log.Fatalf("Server failed to start: %v", newError)
	}
}
