//Guessing game web-app
//Andrew Usevas group-C

package main

import (
	"html/template"
	"net/http"
)

type note struct {
	Message string
}

//SOURCE:http://www.alternatestack.com/development/app-development/simple-http-server-with-templated-response-in-go/
func guessHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("guess.html")               // parses guess.html to create a template "t" (?)internally(?)
	msg := note{Message: "Guess a number between 1 and 20"} // message to pass to templated guess.html
	t.Execute(w, msg)
}

func main() {

	//SOURCES:http://www.alexedwards.net/blog/serving-static-sites-with-go ,
	//https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	fs := http.FileServer(http.Dir("./")) //all files are in the same directory as .go file
	http.Handle("/", fs)

	http.HandleFunc("/guess.html", guessHandler) //invokes this when link "New game" pressed

	http.ListenAndServe(":8080", nil) // port

}
