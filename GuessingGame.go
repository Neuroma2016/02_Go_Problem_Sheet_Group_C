//Guessing game web-app
//Andrew Usevas group-C

package main

import (
	"net/http"
)

func main() {

	//SOURCES:http://www.alexedwards.net/blog/serving-static-sites-with-go ,
	//https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	fs := http.FileServer(http.Dir("./")) //all (web)files are in the same directory as .go file
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil) // port

}
