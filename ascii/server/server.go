package main

import (
	"html/template"
	"log"
	"net/http"

	asciifunc "../functions"
)

var t *template.Template
var tErr *template.Template

func server(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm() //analyse le Formulaire et retourne une erreur (s'il y en a une)
	if err != nil {
		error(w, req, 500)
		return
	}

	//Gestion des erreurs de chemin (avec Error 404)
	if req.URL.Path != "/" && req.URL.Path != "/ascii-art" {
		error(w, req, http.StatusNotFound)
		return
	}
	//Gestion des erreurs de Request (avec Error 400)
	if req.Method != "POST" && req.URL.Path == "/ascii-art" {
		error(w, req, 400)
		return
	} else if req.Method != "GET" && req.URL.Path == "/" {
		error(w, req, 400)
		return
	} else if req.Method != "POST" && req.Method != "GET" {
		error(w, req, 400)
		return
	}

	formSelect := req.Form.Get("select")
	formText := req.Form.Get("text")
	var finalStr string

	if req.URL.Path != "/" {
		finalStr = asciifunc.Ascii(formText, formSelect)

	}

	//Gestion des erreurs interne au serveur (avec les fonctions du programme ASCII-Art) (avec Error 500)
	if finalStr == "" && req.URL.Path == "/ascii-art" {
		error(w, req, 500)
		return
	}

	t.Execute(w, template.HTML(finalStr))

	//fmt.Fprintf(w, "Test Fprintf")

}

func error(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		//fmt.Fprint(w, "Error 404 : Page not found")
		tErr.Execute(w, template.HTML(`<h1>Error 404 : Page not found...<h1>`))
	}

	if status == 400 {
		tErr.Execute(w, template.HTML(`<h1>Error 400 : Bad Request...<h1>`))
	}

	if status == 500 {
		tErr.Execute(w, template.HTML(`<h1>Error 500 : Internal Server Error...<h1>`))
	}
}

func main() {
	t = template.Must(template.ParseFiles("./index.html"))
	tErr = template.Must(template.New("test").Parse("{{.}}"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // récupère tous les fichiers "externe" (comme le style.css)

	http.HandleFunc("/", server) // "/" pour dire qu'on est dans ce chemin et server car cest la fonction

	if err := http.ListenAndServe(":50000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
