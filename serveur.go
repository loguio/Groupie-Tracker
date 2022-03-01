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
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	AddressLocation string `json:"locations"`
	ConcertDatesaddress string `json:"concertDates"`
	RelationsAdress string `json:"relations"`
	Location []string
	ConcertDates []string
	Relations	map[string][]string
	RelationDate [][]string
	DateLocation []DateLocation
}// on créer une structure qui contient toutes les données pouvant etre utile a notre site cela va nous permettre d'afficher chaque groupe avec leur données respectives

type DateLocation struct {
	Location string
	Dates []string
}//Cette structure nous permet d'afficher les lieu ainsi que les dates des spectacles 

type Location struct {
	Id int `json:"id"`
	Location []string `json:"locations"`
	Dates string `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Location

type Dates struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Dates

type Relation struct {
	Id int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
} // cette strcture nous permet de recuperer les donnée du lien API Relation

func clicked(id string) interface{}{
	var oneArtist ArtistAPI		
	var relationdate [][]string
	var nettoyage []DateLocation
	url := "https://groupietrackers.herokuapp.com/api/artists/"+id
	resp, err := http.Get(url) // on recupère les données qui sont stockés dans resp
	if err != nil {
		log.Fatalln(err) // si il y a une erreur donc erreur
	}
	defer resp.Body.Close() 
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
	oneArtist.FirstAlbum = gooddate(oneArtist.FirstAlbum)//on passe la donnée FirstAlbum dans la fonction gooddate pour avoir une date plus explicite
	oneArtist.Location = location(oneArtist.AddressLocation)// on Récupere les données qui nous interesse grace a la fonction Location car AddressLocation est un lien API
	oneArtist.ConcertDates = concertdate(oneArtist.ConcertDatesaddress) // ConcertDatesaddress est aussi un lien API donc on tri les données pur avoir celle qui nous intéresse 
	oneArtist.Relations = relation(oneArtist.RelationsAdress)// on fait la meme chose pour RelationsAdress
	for i:=0;i<len(oneArtist.Location);i++ {
		for k := 0; k< len(oneArtist.Relations[oneArtist.Location[i]]);k++ {
			oneArtist.Relations[oneArtist.Location[i]][k]  = gooddate(oneArtist.Relations[oneArtist.Location[i]][k])//on change aussi la date des concert our avoir une date plus joli
		}
		relationdate=append(relationdate,oneArtist.Relations[oneArtist.Location[i]]) // on rajoute les valeurs des dates dans l'index de la villes correspondante
	}
	oneArtist.RelationDate = relationdate // on stock les valeurs des dates dans 
	oneArtist.DateLocation = nettoyage // on vide notre liste
	for i:=0; i < len(oneArtist.Location);i++ {
		var tempo DateLocation
		tempo.Location = bonLieu(oneArtist.Location[i])
		tempo.Dates = oneArtist.RelationDate[i]
		oneArtist.DateLocation = append(oneArtist.DateLocation,tempo) // on ajoute les valeurs utile dans DateLocation
	}
	return oneArtist
}

func ArtistPage(adress string, Page int) interface{} { //Cette fonction se lance lorsque l'utilisateur est sur la page des artistes 
	fmt.Println("1. Performing Http Get...")
	var idArtist = (Page-1)*12 +1// on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var artists []ArtistAPI// nos artistes seront stockés dans cette variables
	var oneArtist ArtistAPI //on stock les données de un artiste danc cette variable
	var nettoyage []DateLocation
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	for idArtist != Page*12 + 1 { // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
		var relationdate [][]string
		url = "/"+strconv.Itoa(idArtist) // On recupère un URL equivalent a afficher les données avec l'ID d'un artiste
		resp, err := http.Get(adress+url) // on recupère les données qui sont stockés dans resp
		if err != nil {
			log.Fatalln(err) // si il y a une erreur donc erreur
		}
		defer resp.Body.Close() 
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
		idArtist = oneArtist.Id 
		if idArtist == 0 {break} // si l'id est égal a 0 c'est que l'on a atteint la fin des artistes et que il n'y en a pas plus a afficher donc on break pour sortir de la boucle
		oneArtist.FirstAlbum = gooddate(oneArtist.FirstAlbum)//on passe la donnée FirstAlbum dans la fonction gooddate pour avoir une date plus explicite
		oneArtist.Location = location(oneArtist.AddressLocation)// on Récupere les données qui nous interesse grace a la fonction Location car AddressLocation est un lien API
		oneArtist.ConcertDates = concertdate(oneArtist.ConcertDatesaddress) // ConcertDatesaddress est aussi un lien API donc on tri les données pur avoir celle qui nous intéresse 
		oneArtist.Relations = relation(oneArtist.RelationsAdress)// on fait la meme chose pour RelationsAdress
		for i:=0;i<len(oneArtist.Location);i++ {
			for k := 0; k< len(oneArtist.Relations[oneArtist.Location[i]]);k++ {
				oneArtist.Relations[oneArtist.Location[i]][k]  = gooddate(oneArtist.Relations[oneArtist.Location[i]][k])//on change aussi la date des concert our avoir une date plus joli
			}
			relationdate=append(relationdate,oneArtist.Relations[oneArtist.Location[i]]) // on rajoute les valeurs des dates dans l'index de la villes correspondante
		}
		oneArtist.RelationDate = relationdate // on stock les valeurs des dates dans 
		oneArtist.DateLocation = nettoyage // on vide notre liste
		for i:=0; i < len(oneArtist.Location);i++ {
			var tempo DateLocation
			tempo.Location = bonLieu(oneArtist.Location[i])
			tempo.Dates = oneArtist.RelationDate[i]
			oneArtist.DateLocation = append(oneArtist.DateLocation,tempo) // on ajoute les valeurs utile dans DateLocation
		}
		artists = append(artists,oneArtist)// on ajoute un artiste par un dans artists 
		idArtist++// on incremente idArtists pour avoir l'artiste suivant avec l'URL
		fmt.Println(oneArtist)
	}
	return artists // on renvois notre liste avec 12 artistes et les données 
}

func bonLieu(date string)string {
	// cette fonctions permet d'enlever les caractères qui ne servent a rien
	date = strings.Replace(date,"-"," ", -1)
	date = strings.Replace(date,"_"," ", -1)
	return date
}

func relation(adress string) map[string][]string {
	//relation nous permet de prendre les valeurs que l'on a vraiment beosin dans L'API Relation
	var relation Relation
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &relation)
	return relation.DatesLocations
}

func concertdate(adress string) []string{
	// Concertdate nous permet de recuperer les location et les dates dans l'API
	var dates Dates
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &dates)
	for i := 0 ; i< len(dates.Dates);i++ {
		dates.Dates[i] = gooddate(dates.Dates[i])
	}
	return dates.Dates
}

func gooddate(mois string) string {
	// cette fonction nous permet de changer les mois de l'année pour les avoir en texte et on enleve les caractères qui servent a rien
	tempo := strings.Split(mois, "-")
	if tempo[1] == "01" {mois = strings.Replace(tempo[1],"01","Janvier", -1)
	}else if tempo[1] == "02" {mois = strings.Replace(tempo[1],"02","Fevrier", -1)
	}else if tempo[1] == "03" {mois = strings.Replace(tempo[1],"03","Mars", -1)
	}else if tempo[1] == "04" {mois = strings.Replace(tempo[1],"04","Avril", -1)
	}else if tempo[1] == "05" {mois = strings.Replace(tempo[1],"05","Mai", -1)
	}else if tempo[1] == "06" {mois = strings.Replace(tempo[1],"06","Juin", -1)
	}else if tempo[1] == "07" {mois = strings.Replace(tempo[1],"07","Juillet", -1)
	}else if tempo[1] == "08" {mois = strings.Replace(tempo[1],"08","Aout", -1)
	}else if tempo[1] == "09" {mois = strings.Replace(tempo[1],"09","Septembre", -1)
	}else if tempo[1] == "10" {mois = strings.Replace(tempo[1],"10","Octobre", -1)
	}else if tempo[1] == "11" {mois = strings.Replace(tempo[1],"11","Novembre", -1)
	}else if tempo[1] == "12" {mois = strings.Replace(tempo[1],"12","Decembre", -1)}
	tempo[1] = mois
	mois = strings.Join(tempo, " ")
	mois = strings.Replace(mois,"*","",-1)
	return mois
}

func location(adress string) []string {
	// on recupere les information qui nous interresse dans locations
	var locations Location
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &locations)
	return locations.Location
}

func main() {
	// fmt.Println(clicked("1"))
	lien := "https://groupietrackers.herokuapp.com/api"
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	// affiche l'html
	tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
	if err != nil {
	}
	page := 1

	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		data := ArtistPage(lien+"/artists", page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/PageSuivante", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/navPage.gohtml")
		}
		page+=1
		data := ArtistPage(lien+"/artists",page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/PagePrecedente", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/navPage.gohtml")
		}
		page-=1
		data := ArtistPage(lien+"/artists",page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}