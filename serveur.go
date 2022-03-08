package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
	http.HandleFunc("/Groupie-tracker/nbArtist", nbArtist)
	http.HandleFunc("/Groupie-tracker/artist", artist)
	http.HandleFunc("/Groupie-tracker/Recherche", rechercher)
	http.HandleFunc("/Groupie-tracker/listartistA-Z", FiltreAlpha)
	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

//#####################################################################################################################################//

func rechercher(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		data, errr := rechercheFind("https://groupietrackers.herokuapp.com/api/artists", strings.Split(strings.ToUpper(recherche), ""))
		if errr != nil {
			print(errr)
		}
		tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
		if err != nil {
			print(err)
		}
		tmpl.ExecuteTemplate(w, "listartists", data)
	}
}

//##########################################################################################################################################//

func groupieTracker(w http.ResponseWriter, r *http.Request) { //récupération des donnée a envoyer sur la page html
	tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		fmt.Println(err)
		tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		print(err)
	}
	tmpl.ExecuteTemplate(w, "home", "") //exécution du template
}

//#######################################################################################################################################//

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

//#########################################################################################################################################//

func nbArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		page := 1
		lien := "https://groupietrackers.herokuapp.com/api"
		nbArtist, errr := strconv.Atoi(r.FormValue("Artists"))
		tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
		if errr != nil || err != nil {
		}
		fmt.Println(page, nbArtist)
		data, err := ArtistPage(lien+"/artists", page, nbArtist) //récupération des donnée a envoyer sur la page html
		tmpl.ExecuteTemplate(w, "listartists", data)             //exécution du template
	}
}

//##############################################################################################################################//

func listartist(w http.ResponseWriter, r *http.Request) {
	lien := "https://groupietrackers.herokuapp.com/api"
	page := 1
	nbArtist := 12
	tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		fmt.Println(err, "/")
		tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		if err != nil {
			fmt.Println(err)
		}
	}
	data, err := ArtistPage(lien+"/artists", page, nbArtist)
	if err != nil {
		fmt.Println(err, "/")
		tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
		fmt.Println(err)
	} //récupération des donnée a envoyer sur la page html
	tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
}

//########################################################################################################################################//

func PageSuivante(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		lien := "https://groupietrackers.herokuapp.com/api"
		page, err := strconv.Atoi(r.FormValue("page"))
		nbArtist, errr := strconv.Atoi(r.FormValue("Artists"))
		if errr != nil {
		}
		if err != nil {
			fmt.Println("erreur page")
			tmpl, err := template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			if err != nil {
				fmt.Println("error 501")
			}
			data, err := ArtistPage(lien+"/artists", page, nbArtist) //récupération des donnée a envoyer sur la page html
			if err != nil {
				fmt.Println("error data")
			}
			tmpl.ExecuteTemplate(w, "index", data) //exécution du template
			return
		}
		tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
		if err != nil {
			fmt.Println(err)
			tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			print(err)
		}
		page += 1
		data, err := ArtistPage(lien+"/artists", page, nbArtist) //récupération des donnée a envoyer sur la page html
		if err != nil {
			tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			fmt.Println(err)
		}
		tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
		return

	}
}

//#########################################################################################################################################//

func PagePrecedente(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		lien := "https://groupietrackers.herokuapp.com/api"
		page, err := strconv.Atoi(r.FormValue("page"))
		nbArtist, errr := strconv.Atoi(r.FormValue("Artists"))
		if errr != nil {
		}
		if err != nil {
			fmt.Println("erreur page")
			tmpl, err := template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			if err != nil {
				fmt.Println(err)
			}
			data, err := ArtistPage(lien+"/artists", page, nbArtist) //récupération des donnée a envoyer sur la page html
			tmpl.ExecuteTemplate(w, "index", data)                   //exécution du template
			return
		}
		tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
		if err != nil {
			fmt.Println(err, "UWU")
			tmpl, err = template.ParseFiles("./templates/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			print(err)
		}
		if page < 1 {
			page -= 1
		}
		data, err := ArtistPage(lien+"/artists", page, nbArtist) //récupération des donnée a envoyer sur la page html
		if err != nil {
			tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			if err != nil {
				fmt.Println(err)
			}
		}
		tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
		return
	}
}

//#############################################################################################################################################//

func FiltreAlpha(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		data, err := trie(lien+"/artists", page)
		if err != nil {
			fmt.Println(err, "/")
			tmpl, err = template.ParseFiles("./assets/Error500.gohtml") //utilisation du fichier Error500.gohtml pour le template
			fmt.Println(err)
		} //récupération des donnée a envoyer sur la page html
		tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
	}
}
