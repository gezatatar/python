package main

import (
	"fmt"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World!\n")
}

func main() {
	//fmt.Println("Hello world!")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":8080", nil)
}
