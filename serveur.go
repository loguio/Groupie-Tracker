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

type Page struct {
	Name         string
	Image 		 string
	Members      []string
	CreationDate string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
	Page int
}

type ArtistAPI struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Location string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations string `json:"relations"`
	Page int
}

func HomePage(adress string, nbPage int) interface{} {
	fmt.Println("1. Performing Http Get...")
	var idArtist = (nbPage-1)*12 +1
	var url = ""
	var artists []ArtistAPI
	var oneartist ArtistAPI
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
		artists = append(artists,oneartist)
		if idArtist == 0 {
			break
		}
		idArtist++
	}
	return artists
}


func get(adress string, nbArtist int) (interface{}, int) {
	valreturn := 0
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
		valreturn = 500
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var tab interface{}
	// Convert response body to tab interface
	json.Unmarshal(bodyBytes, &tab)
	var name string
	var image string
	var members []string
	var creationDate string
	var firstAlbum string
	var location string
	var concertDates string
	var relations string
	var page = 0
	for key, value := range tab.([]interface{})[nbArtist].(map[string]interface{}) {
		if key == "name" {
			name = fmt.Sprint(value)
		}
		if key == "members" {
			members = strings.Split(fmt.Sprint(value), " ")
		}
		if key == "creationDate" {
			creationDate = fmt.Sprint(value)
		}
		if key == "firstAlbum" {
			firstAlbum = fmt.Sprint(value)
		}
		if key == "locations" {
			location = fmt.Sprint(value)
		}
		if key == "concertDates" {
			concertDates = fmt.Sprint(value)
		}
		if key == "relations" {
			relations = fmt.Sprint(value)
		}
		if key == "image" {
			image = fmt.Sprint(value)
		}
	}
	data := Page{name, image,members, creationDate, firstAlbum, location, concertDates, relations, page}
	return data,valreturn
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	// affiche l'html
	tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
	if err != nil {
	}
	nb := 4
	page := 1

	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			nb, _ = strconv.Atoi(r.FormValue("nombre"))
		}
		data, codeError := get(lien+"/artists", nb)
		if codeError == 500 {

		}
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

	http.HandleFunc("/Groupie-tracker/artist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/navPage.gohtml")
			nb, _ = strconv.Atoi(r.FormValue("nombre"))
		}
		data := HomePage(lien+"/artists",page)
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
