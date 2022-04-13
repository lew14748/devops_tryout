package main
 
import (
"fmt"
"net/http"
)
 
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "nothing interesting here")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, stranger")
	})
	http.HandleFunc("/hello_there", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "General Kenobi?")
	})
	http.ListenAndServe(":2345", nil)
}
