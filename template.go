package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAgg struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func templateHandler(w http.ResponseWriter, r *http.Request) {

	p := NewsAgg{"Some title", "Some News"}
	t, _ := template.ParseFiles("basictemplate.html")

	error := t.Execute(w, p)

	if error != nil {
		fmt.Fprintf(w, error.Error())
	}

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/yo", templateHandler)
	http.ListenAndServe(":8000", nil)
}
