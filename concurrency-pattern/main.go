package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Fprintf(w, "GET Method")

	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "POST Method: %s", body)
	}
}
func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
