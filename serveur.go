package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//on Importe toute les bibliothèques que l'on a besoin
type ArtistAPI struct {
	Id                  int      `json:"id"`
	Image               string   `json:"image"`
	Name                string   `json:"name"`
	Members             []string `json:"members"`
	CreationDate        int      `json:"creationDate"`
	FirstAlbum          string   `json:"firstAlbum"`
	AddressLocation     string   `json:"locations"`
	ConcertDatesaddress string   `json:"concertDates"`
	RelationsAdress     string   `json:"relations"`
	Location            []string
	ConcertDates        []string
	Relations           map[string][]string
	RelationDate        [][]string
	DateLocation        []DateLocation
} // on créer une structure qui contient toutes les données pouvant etre utile a notre site cela va nous permettre d'afficher chaque groupe avec leur données respectives
type DateLocation struct {
	Location string
	Dates    []string
} //Cette structure nous permet d'afficher les lieu ainsi que les dates des spectacles
type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Location
type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Dates
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
} // cette strcture nous permet de recuperer les donnée du lien API Relation
//on Importe toute les bibliothèques que l'on a besoin
func clicked(id string) (interface{}, error) {
	var oneArtist ArtistAPI
	var relationdate [][]string
	var clean []DateLocation
	url := "https://groupietrackers.herokuapp.com/api/artists/" + id
	resp, err := http.Get(url) // on recupère les données qui sont stockés dans resp
	if err != nil {
		log.Fatalln(err) // si il y a une erreur donc erreur
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &oneArtist)                         // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
	oneArtist.FirstAlbum = gooddate(oneArtist.FirstAlbum)         //on passe la donnée FirstAlbum dans la fonction gooddate pour avoir une date plus explicite
	oneArtist.Location, err = location(oneArtist.AddressLocation) // on Récupere les données qui nous interesse grace a la fonction Location car AddressLocation est un lien API
	if err != nil {
		return oneArtist, err
	}
	oneArtist.ConcertDates, err = concertdate(oneArtist.ConcertDatesaddress) // ConcertDatesaddress est aussi un lien API donc on tri les données pur avoir celle qui nous intéresse
	if err != nil {
		return oneArtist, err
	}
	oneArtist.Relations, err = relation(oneArtist.RelationsAdress) // on fait la meme chose pour RelationsAdress
	if err != nil {
		return oneArtist, err
	}
	for i := 0; i < len(oneArtist.Location); i++ {
		for k := 0; k < len(oneArtist.Relations[oneArtist.Location[i]]); k++ {
			oneArtist.Relations[oneArtist.Location[i]][k] = gooddate(oneArtist.Relations[oneArtist.Location[i]][k]) //on change aussi la date des concert our avoir une date plus joli
		}
		relationdate = append(relationdate, oneArtist.Relations[oneArtist.Location[i]]) // on rajoute les valeurs des dates dans l'index de la villes correspondante
	}
	oneArtist.RelationDate = relationdate // on stock les valeurs des dates dans
	oneArtist.DateLocation = clean        // on vide notre liste
	for i := 0; i < len(oneArtist.Location); i++ {
		var tempo DateLocation
		tempo.Location = bonLieu(oneArtist.Location[i])
		tempo.Dates = oneArtist.RelationDate[i]
		oneArtist.DateLocation = append(oneArtist.DateLocation, tempo) // on ajoute les valeurs utile dans DateLocation
	}
	return oneArtist, err
}

func ArtistPage(adress string, Page int) (interface{}, error) { //Cette fonction se lance lorsque l'utilisateur est sur la page des artistes
	fmt.Println("1. Performing Http Get...")
	var idArtist = (Page-1)*12 + 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var artists []ArtistAPI // nos artistes seront stockés dans cette variables
	var oneArtist ArtistAPI //on stock les données de un artiste danc cette variable
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	for idArtist != Page*12+1 { // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
		url = "/" + strconv.Itoa(idArtist)  // On recupère un URL equivalent a afficher les données avec l'ID d'un artiste
		resp, err := http.Get(adress + url) // on recupère les données qui sont stockés dans resp
		if err != nil {
			fmt.Println(err) // si il y a une erreur donc erreur
			return artists, err
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
		idArtist = oneArtist.Id
		if idArtist == 0 {
			fmt.Println("erreur : L'API est vide")
			break
		} // si l'id est égal a 0 c'est que l'on a atteint la fin des artistes et que il n'y en a pas plus a afficher donc on return pour sortir de la boucle
		oneArtist.FirstAlbum = gooddate(oneArtist.FirstAlbum)         //on passe la donnée FirstAlbum dans la fonction gooddate pour avoir une date plus explicite
		oneArtist.Location, err = location(oneArtist.AddressLocation) // on Récupere les données qui nous interesse grace a la fonction Location car AddressLocation est un lien API
		if err != nil {
			return artists, err
		}
		artists = append(artists, oneArtist)
		idArtist++
	}
	var err error
	return artists, err // on renvois notre liste avec 12 artistes et les données
}

func bonLieu(date string) string {
	// cette fonctions permet d'enlever les caractères qui ne servent a rien
	date = strings.Replace(date, "-", " ", -1)
	date = strings.Replace(date, "_", " ", -1)
	return date
}

func relation(adress string) (map[string][]string, error) {
	//relation nous permet de prendre les valeurs que l'on a vraiment beosin dans L'API Relation
	var relation Relation
	resp, err := http.Get(adress)
	if err != nil {
		fmt.Println(err)
		return relation.DatesLocations, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &relation)
	return relation.DatesLocations, err
}

func concertdate(adress string) ([]string, error) {
	// Concertdate nous permet de recuperer les location et les dates dans l'API
	var dates Dates
	resp, err := http.Get(adress)
	if err != nil {
		fmt.Println(err)
		return dates.Dates, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &dates)
	for i := 0; i < len(dates.Dates); i++ {
		dates.Dates[i] = gooddate(dates.Dates[i])
	}
	return dates.Dates, err
}

func gooddate(mois string) string {
	// cette fonction nous permet de changer les mois de l'année pour les avoir en texte et on enleve les caractères qui servent a rien
	tempo := strings.Split(mois, "-")
	if tempo[1] == "01" {
		mois = strings.Replace(tempo[1], "01", "Janvier", -1)
	} else if tempo[1] == "02" {
		mois = strings.Replace(tempo[1], "02", "Fevrier", -1)
	} else if tempo[1] == "03" {
		mois = strings.Replace(tempo[1], "03", "Mars", -1)
	} else if tempo[1] == "04" {
		mois = strings.Replace(tempo[1], "04", "Avril", -1)
	} else if tempo[1] == "05" {
		mois = strings.Replace(tempo[1], "05", "Mai", -1)
	} else if tempo[1] == "06" {
		mois = strings.Replace(tempo[1], "06", "Juin", -1)
	} else if tempo[1] == "07" {
		mois = strings.Replace(tempo[1], "07", "Juillet", -1)
	} else if tempo[1] == "08" {
		mois = strings.Replace(tempo[1], "08", "Aout", -1)
	} else if tempo[1] == "09" {
		mois = strings.Replace(tempo[1], "09", "Septembre", -1)
	} else if tempo[1] == "10" {
		mois = strings.Replace(tempo[1], "10", "Octobre", -1)
	} else if tempo[1] == "11" {
		mois = strings.Replace(tempo[1], "11", "Novembre", -1)
	} else if tempo[1] == "12" {
		mois = strings.Replace(tempo[1], "12", "Decembre", -1)
	}
	tempo[1] = mois
	mois = strings.Join(tempo, " ")
	mois = strings.Replace(mois, "*", "", -1)
	return mois
}

func location(adress string) ([]string, error) {
	// on recupere les information qui nous interresse dans locations
	var locations Location
	resp, err := http.Get(adress)
	if err != nil {
		fmt.Println(err)
		return locations.Location, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &locations)
	for i := 0; i < len(locations.Location); i++ {
		for k := 0; k < len(locations.Location); k++ {
			if locations.Location[i] == locations.Location[k] && i != k {
				locations.Location[i] = ""
			}
		}
	}
	return locations.Location, err
}

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
