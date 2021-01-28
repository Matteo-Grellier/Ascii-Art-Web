package main

import (
	"html/template"
	"log"
	"net/http"

	asciifunc "../templates/functions"
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

	formSelect := req.Form.Get("select") //récupère l'information du formulaire (au nom "select")
	formText := req.Form.Get("text")     //récupère l'information du formulaire (au nom "text")
	var finalStr string

	if req.URL.Path != "/" { //Vérification pour éviter que le programme "ascii.go" se lance dès le début (évite une erreur 500)
		finalStr = asciifunc.Ascii(formText, formSelect)
	}

	//Gestion des erreurs interne au serveur (avec les fonctions du programme ASCII-Art) (avec Error 500)
	if finalStr == "" && req.URL.Path == "/ascii-art" {
		error(w, req, 500)
		return
	}

	t.Execute(w, template.HTML(finalStr)) // "t.Execute" va écrire une réponse à la requête et "template.HTML" va écrire du texte sous forme HTML à l'emplacement "{{.}}" sur la page HTML.

}

func error(w http.ResponseWriter, r *http.Request, errorCode int) {
	w.WriteHeader(errorCode)

	//Si le code erreur est 404 (404 = http.StatusNotFound)
	if errorCode == http.StatusNotFound {
		tErr.Execute(w, template.HTML(`<h1>Error 404 : Page not found...<h1>`)) //(voir ligne 53 pour explication)
	}

	//Si le code erreur est 400
	if errorCode == 400 {
		tErr.Execute(w, template.HTML(`<h1>Error 400 : Bad Request...<h1>`)) //(voir ligne 53 pour explication)
	}

	//Si le code erreur est 500
	if errorCode == 500 {
		tErr.Execute(w, template.HTML(`<h1>Error 500 : Internal Server Error...<h1>`)) //(voir ligne 53 pour explication)
	}
}

func main() {
	t = template.Must(template.ParseFiles("./index.html"))        // Récupère la page "index.html" comme page du template
	tErr = template.Must(template.New("forError").Parse("{{.}}")) // Créer une page du template dans lequel il y a "{{.}}"

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))          // récupère tous les fichiers "externe" dans "css"
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("../images")))) // récupère tous les fichiers "externe" dans "images"

	http.HandleFunc("/", server) // Permet d'utiliser une fonction pour répondre à la requête HTTP (fonction "server") dans le chemin "/"

	if err := http.ListenAndServe(":50000", nil); err != nil { // "ListenAndServe" écoute les requêtes HTTP (ici sur le port 50000)
		log.Fatal("ListenAndServe: ", err)
	}
}
