//Guessing game web-app Starter
//Andrew Usevas group-C
//SOURCE:https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/03.2.html

package main

import (
	"fmt"
	"log"
	"net/http"
)

func reqHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Guessing game") //will be displayed in the body
}

func main() {
	http.HandleFunc("/", reqHandler)
	err := http.ListenAndServe(":8080", nil) // port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
