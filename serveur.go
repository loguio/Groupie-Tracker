package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func main() {
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// affiche l'html
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	data := Page{"PLZZZZZZZZZZZZ"}
	if err != nil {
	}
	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", data)
	})

	/////////////////////////////////

	fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur

}
