package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("renderHomePage")
	fmt.Fprint(w, "<p>Hello from Docker</p>")
}

func renderAboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("renderAboutPage")
	fmt.Fprint(w, "<p>Hello about page</p>")
}

func main() {
	fmt.Println("hello world")
	router := mux.NewRouter()

	router.HandleFunc("/", renderHomePage).Methods("POST")
	// http.HandleFunc("/", renderHomePage)
	// http.HandleFunc("/about", renderAboutPage)
	http.ListenAndServe(":3000", router)
}
