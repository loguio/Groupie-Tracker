package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//on Importe toute les bibliothèques que l'on a besoin

func main() {
	fileServer := http.FileServer(http.Dir("static/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// affiche l'html
	http.HandleFunc("/Groupie-tracker", groupieTracker)
	http.HandleFunc("/Groupie-tracker/PageSuivante", PageSuivante)
	http.HandleFunc("/Groupie-tracker/PagePrecedente", PagePrecedente)
	http.HandleFunc("/Groupie-tracker/listartist", listartist)
	http.HandleFunc("/Groupie-tracker/artist", artist)
	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

func groupieTracker(w http.ResponseWriter, r *http.Request) {
	lien := "https://groupietrackers.herokuapp.com/api"
	page := 1
	data, err := ArtistPage(lien+"/artists", page)                                                                                                                                      //récupération des donnée a envoyer sur la page html
	tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		fmt.Println(err)
		tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		print(err)
	}
	tmpl.ExecuteTemplate(w, "home", data) //exécution du template
}

func artist(w http.ResponseWriter, r *http.Request) {
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
}

func listartist(w http.ResponseWriter, r *http.Request) {
	lien := "https://groupietrackers.herokuapp.com/api"
	page := 1
	tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		fmt.Println(err, "/")
		tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		if err != nil {
			fmt.Println(err)
		}
	}
	data, err := ArtistPage(lien+"/artists", page)
	if err != nil {
		fmt.Println(err, "/")
		tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		fmt.Println(err)
	} //récupération des donnée a envoyer sur la page html
	tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
}

func PageSuivante(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.URL.Path != "/Groupie-tracker/PageSuivante" {
			errorHandler(w, r, http.StatusNotFound)
			return
		} else {
			lien := "https://groupietrackers.herokuapp.com/api"
			page, err := strconv.Atoi(r.FormValue("page"))
			if err != nil {
				fmt.Println("erreur page")
				tmpl, err := template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				if err != nil {
					fmt.Println("error 501")
				}
				data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
				if err != nil {
					fmt.Println("error data")
				}
				tmpl.ExecuteTemplate(w, "index", data) //exécution du template
				return
			}
			tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				fmt.Println(err, "UWU")
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			page += 1
			fmt.Println(page)
			data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
			if err != nil {
				tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			}
			tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
			return
		}
	}
}

func PagePrecedente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.URL.Path != "/Groupie-tracker/PagePrecedente" {
			errorHandler(w, r, http.StatusNotFound)
			return
		} else {
			lien := "https://groupietrackers.herokuapp.com/api"
			page, err := strconv.Atoi(r.FormValue("page"))
			if err != nil {
				fmt.Println("erreur page")
				tmpl, err := template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				if err != nil {
				}
				data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
				tmpl.ExecuteTemplate(w, "index", data)         //exécution du template
				return
			}
			tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				fmt.Println(err, "UWU")
				tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
				print(err)
			}
			page -= 1
			fmt.Println(page)
			data, err := ArtistPage(lien+"/artists", page) //récupération des donnée a envoyer sur la page html
			if err != nil {
				tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			}
			tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
			return
		}
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
