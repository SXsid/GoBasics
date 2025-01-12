package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "Post req successful\n")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Fprintf(w, "Email is%s\n", email)
	fmt.Fprintf(w, "Pass is%s\n ", password)
}
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/get" {
		http.Error(w, "404,not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "404,not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello from the sever")
}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("server is up on :8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
