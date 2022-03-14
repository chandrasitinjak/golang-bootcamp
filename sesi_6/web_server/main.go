package main

import (
	"fmt"
	"net/http"
)

//var PORT = ":8000"

func main() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello worldss"
	//fmt.Println("hai hai", w, msg)
	fmt.Fprint(w, msg)
}
