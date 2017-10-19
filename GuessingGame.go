//Guessing game web-app
//Andrew Usevas group-C

package main

import (
	"fmt"
	"log"
	"net/http"
)

func reqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html") //will render <h1> below as html
	fmt.Fprintf(w, "<h1>Guessing game</h>")     //will be displayed in the body
}

func main() {
	http.HandleFunc("/", reqHandler)
	err := http.ListenAndServe(":8080", nil) // port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
