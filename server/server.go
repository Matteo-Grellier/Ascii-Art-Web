package main

import (
	"html/template"
	"net/http"
)

type NameData struct {
	FirstName string
	LastName  string
}

var t *template.Template

func home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(400)
	t.Execute(w, NameData{
		FirstName: "Matteo",
		LastName:  "Grellier",
	})
}

func main() {
	t = template.Must(template.ParseFiles("./index.html"))
	http.HandleFunc("/", home)
	http.HandleFunc("/ascii-art", home)
	http.ListenAndServe(":8000", nil)
}
