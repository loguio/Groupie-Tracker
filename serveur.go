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
	http.HandleFunc("/", error404)                             // lance l'erreur 404 quand on est sur une URL pas utilisée
	http.HandleFunc("/Groupie-tracker", groupieTracker)        // lance la fonction Groupie tracket sur l'url "groupie-tracker"
	http.HandleFunc("/Groupie-tracker/listartist", listartist) //lance la fonction listartists sur cette url
	http.HandleFunc("/Groupie-tracker/nbArtist", nbArtist)     // lance la fonction nbartists sur cette url
	http.HandleFunc("/Groupie-tracker/artist", artist)         // lance la fonction artists sur l'url "artists"
	http.HandleFunc("/Groupie-tracker/Recherche", rechercher)  // lance la fonction Find sur l'url "find"
	http.HandleFunc("/Groupie-tracker/cart", carte)            // lance la fonction Carte sur l'url "carte"
	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}

//#####################################################################################################################################//

func error404(w http.ResponseWriter, r *http.Request) { // fonction qui affiche la page de l'erreur 404
	var data interface{}
	tmpl, err := template.ParseFiles("./templates/Error404.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "error404", data) // exécute le template sur la page html
}

//########################################################################################################################################//

func rechercher(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		recherche := r.FormValue("recherche")
		data, errr := rechercheFind("https://groupietrackers.herokuapp.com/api/artists", strings.Split(strings.ToUpper(recherche), "")) // récupération des artiste pour les artists recherché
		if errr != nil {
			error500(errr, w)
		} else {
			tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				error500(errr, w)
			} else {
				tmpl.ExecuteTemplate(w, "listartists", data)
			}
		}
	}
}

//##########################################################################################################################################//

func groupieTracker(w http.ResponseWriter, r *http.Request) { //récupération des donnée a envoyer sur la page html
	tmpl, err := template.ParseFiles("./templates/home.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageaccueil.html", "./templates/pagelistartists.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		error500(err, w)
	}
	tmpl.ExecuteTemplate(w, "home", "") //exécution du template
}

//#######################################################################################################################################//

func artist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id_artiste := r.FormValue("id")
		data, err := clicked(id_artiste) //récupération des donnée a envoyer sur la page html
		if err != nil {
			error500(err, w)
		} else {
			tmpl, err := template.ParseFiles("./templates/artist.html", "./templates/navbar.html", "./templates/footer.html", "./templates/pageartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				error500(err, w)
			} else {
				tmpl.ExecuteTemplate(w, "artist", data) //exécution du template
			}
		}
	}
}

//#########################################################################################################################################//

func nbArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		page := 1
		lien := "https://groupietrackers.herokuapp.com/api"
		nbArtist, err := strconv.Atoi(r.FormValue("Artists"))
		if err != nil {
			error500(err, w)
		} else {
			tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
			if err != nil {
				error500(err, w)
			} else {
				function := r.FormValue("function")
				data, err := ArtistPage(lien+"/artists", page, nbArtist, function) //récupération des donnée a envoyer sur la page html
				if err != nil {
					error500(err, w)
				} else {
					tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
				}
			}
		}
	}
}

//####################################################################################################################################

func carte(w http.ResponseWriter, r *http.Request) {
	var Page Carte
	lien := "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q="
	tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html", "./templates/pagecart.html")
	if err != nil {
		error500(err, w)
	} else {
		value := r.FormValue("carte")
		if value == "" {
			lien = "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q=Paris"
		}
		Page.Location = lieux("https://groupietrackers.herokuapp.com/api/locations/")
		Page.Valeur = lien + value
		tmpl.ExecuteTemplate(w, "pagecart", Page)
	}
}
