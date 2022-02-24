package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Name         string
	Image 		 string
	Members      []string
	CreationDate string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

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
	Relations []string
}

type Location struct {
	Id int `json:"id"`
	Location []string `json:"locations"`
	Dates string `json:"dates"`
}

type Dates struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id int `json:"id"`
	// DatesLocations {}interface `json:"datesLocations"`
}

func HomePage(adress string, nbPage int) interface{} {
	fmt.Println("1. Performing Http Get...")
	var idArtist = (nbPage-1)*12 +1
	var url = ""
	var artists []ArtistAPI
	var oneartist ArtistAPI
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	for idArtist != nbPage*12 + 1 {
		url = "/"+strconv.Itoa(idArtist)
		resp, err := http.Get(adress+url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneartist)
		idArtist = oneartist.Id 
		if idArtist == 0 {
			break
		}
		oneartist.Location = location(oneartist.AddressLocation)
		oneartist.ConcertDates = concertdate(oneartist.ConcertDatesaddress)
		artists = append(artists,oneartist)
		idArtist++
	}
	fmt.Println(artists)
	return artists
}
func concertdate(adress string) []string{
	var dates Dates
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &dates)
	fmt.Println(dates)
	return dates.Dates
}

func location(adress string) []string {
	var location Location
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &location)
	fmt.Println(location)
	return location.Location
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	// affiche l'html
	tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
	if err != nil {
	}
	page := 1

	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		data := HomePage(lien+"/artists", page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/PageSuivante", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/navPage.gohtml")
		}
		page+=1
		data := HomePage(lien+"/artists",page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/PagePrecedente", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/navPage.gohtml")
		}
		page-=1
		data := HomePage(lien+"/artists",page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
