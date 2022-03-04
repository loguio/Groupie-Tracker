package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//on Importe toute les bibliothèques que l'on a besoin

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	fileServer := http.FileServer(http.Dir("static/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// affiche l'html
	page := 1

	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) { //récupération des donnée a envoyer sur la page html
		tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html") // utilisation du fichier navPage.gohtml pour le template
		if err != nil {
			fmt.Println(err, "error /groupie-tracker")
			tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			print(err)
		}
		tmpl.ExecuteTemplate(w, "home", "data") //exécution du template
	})

	http.HandleFunc("/Groupie-tracker/artist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id_artiste := r.FormValue("id")
			data, err := clicked(id_artiste) //récupération des donnée a envoyer sur la page html
			if err != nil {
				fmt.Println(err, "UWU")
			}
			tmpl, err := template.ParseFiles("./templates/artist.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				fmt.Println(err, "UWU")
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			tmpl.ExecuteTemplate(w, "artist", data) //exécution du template
		}
	})

	http.HandleFunc("/Groupie-tracker/listartist", func(w http.ResponseWriter, r *http.Request) {
		data, err := ArtistPage(lien+"/artists", page)
		print(err)
		tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template // utilisation du fichier navPage.gohtml pour le template
		if err != nil {
			fmt.Println(err, "/")
			tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			print(err)
		}
		tmpl.ExecuteTemplate(w, "listartist", data) //exécution du template
	})

	http.HandleFunc("/Groupie-tracker/PageSuivante", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template// utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				fmt.Println(err, "UWU")
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			page += 1
			data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
			if err != nil {
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			tmpl.ExecuteTemplate(w, "listartist", data) //exécution du template
		}
	})

	http.HandleFunc("/Groupie-tracker/PagePrecedente", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				fmt.Println(err, "UWU")
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			page -= 1
			data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
			if err != nil {
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			tmpl.ExecuteTemplate(w, "listartist", data) //exécution du template
		}

	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
