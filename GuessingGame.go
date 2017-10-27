//Guessing game web-app
//Andrew Usevas group-C

package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type note struct {
	Message string
}

//SOURCE: https://golangcode.com/generate-random-numbers/
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//SOURCE:http://www.alternatestack.com/development/app-development/simple-http-server-with-templated-response-in-go/
//https://github.com/data-representation/go-cookies/blob/master/go-cookie.go
func guessHandler(w http.ResponseWriter, r *http.Request) {
	count := 0
	rand.Seed(time.Now().UnixNano())
	count = random(1, 20)

	// Try to read the cookie.
	var cookie, err = r.Cookie("target")
	if err == nil {
		// If we could read it, try to convert its value to an int.
		count, _ = strconv.Atoi(cookie.Value)
	}

	// Create a cookie instance and set the cookie.
	cookie = &http.Cookie{
		Name:  "target",
		Value: strconv.Itoa(count),
	}
	http.SetCookie(w, cookie)

	t, _ := template.ParseFiles("guess.html")               // parses guess.html to create a template "t" (?)internally(?)
	msg := note{Message: "Guess a number between 1 and 20"} // message to pass to templated guess.html
	t.Execute(w, msg)
}

func main() {

	//SOURCES:http://www.alexedwards.net/blog/serving-static-sites-with-go ,
	//https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	fs := http.FileServer(http.Dir("./")) //all files are in the same directory as .go file
	http.Handle("/", fs)

	http.HandleFunc("/guess.html", guessHandler) //invoked when "New game" pressed

	http.ListenAndServe(":8080", nil) // port

}
