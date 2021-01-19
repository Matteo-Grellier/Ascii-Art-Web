package main

import (
	"fmt"
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

	//Gestion des erreurs de chemin (avec Error 404)
	if req.URL.Path != "/" && req.URL.Path != "/ascii-art" {
		errorHandler(w, req, http.StatusNotFound)
		return
	}

	formSelect := req.Form.Get("select")
	formText := req.Form.Get("text")

	finalStr := asciifunc.Ascii(formText, formSelect)

	//fmt.Fprintf(w, "Test Fprintf")

	t.Execute(w, template.HTML(finalStr))

}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error 404 : Page not found")
	}
}

func main() {
	t = template.Must(template.ParseFiles("./index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // récupère tous les fichiers "externe" (comme le style.css)

	http.HandleFunc("/", server) // "/" pour dire qu'on est dans ce chemin et server car cest la fonction

	http.ListenAndServe(":50000", nil)
}
