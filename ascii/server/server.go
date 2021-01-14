package main

import (
	"html/template"
	"net/http"

	asciifunc "../functions"
)

var t *template.Template

func server(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() //cette ligne est (je crois) indispensable.
	if err != nil {
		panic(err)
	}

	formSelect := req.Form.Get("select")
	formText := req.Form.Get("text")

	finalStr := asciifunc.Ascii(formText, formSelect)

	t.Execute(w, template.HTML(finalStr))
}

func main() {
	t = template.Must(template.ParseFiles("./index.html"))

	// fs := http.FileServer(http.Dir("../style"))
	// http.Handle("style", http.StripPrefix("style", fs))

	http.Handle("style", http.FileServer(http.Dir("../style")))

	http.HandleFunc("/", server) // "/" pour dire qu'on est dans ce fichier (je crois) et server car cest la fonction
	//http.Handle("/", fs)
	// http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("../style"))))
	http.ListenAndServe(":50000", nil)
}
