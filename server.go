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
	http.HandleFunc("/Groupie-tracker/Recherche", find)        // lance la fonction Find sur l'url "find"
	http.HandleFunc("/Groupie-tracker/cart", mapp)             // lance la fonction Carte sur l'url "carte"
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

func find(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // on recupère l'information que l'utilisateur a bien rentrer une information
		find := r.FormValue("find")                                                                                       // on recupère la veleur et on la stock dans find
		data, errr := Find("https://groupietrackers.herokuapp.com/api/artists", strings.Split(strings.ToUpper(find), "")) //on recupère les artistes qui on les bons artistes par rapport a find
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
	if r.Method == "POST" { // on recupère l'information que l'utilisateur a bien rentrer une information
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

func mapp(w http.ResponseWriter, r *http.Request) {
	var Page Carte                                                                                   // on créer page qui va nous permette d'avoir toute les donnéés utiles
	url := "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q=" // cette url nous permet d'avoir la map de google maps avec la clé de l'API
	tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/pagecart.html")
	if err != nil {
		error500(err, w)
	}
	value := r.FormValue("map") // on recupère la valeur map sur le serveur ce qui va nous permette d'avoir le lieu que l'utilisateur a demander
	if value == "" {
		url = "https:www.google.com/maps/embed/v1/place?key=AIzaSyAXXPpGp3CYZDcUSiE2YRlNID4ybzoZa7o&q=Paris" // lorsque l'on lance la map pour la première fois c'est Paris qui est affiché
	}
	Page.Location = place("https://groupietrackers.herokuapp.com/api/locations/") // on envois toute les lieux possible sans doublons et dans l'ordre alphabetique a l'utilisateur
	Page.Valeur = url + value
	tmpl.ExecuteTemplate(w, "pagecart", Page) // on affiche le template pagecart avec les données de Page
}
