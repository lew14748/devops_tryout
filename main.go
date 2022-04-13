package main
 
import (
"fmt"
"net/http"
)
 
func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "nothing interesting here")
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, stranger")
}
func helloThereHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "General Kenobi?")
}
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello_there", helloThereHandler)
	http.ListenAndServe(":2345", nil)
}
