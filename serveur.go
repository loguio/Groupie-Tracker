package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Name         string
	Members      string
	CreationDate string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

func get(adress string) interface{} {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var tab interface{}
	// Convert response body to tab interface
	json.Unmarshal(bodyBytes, &tab)
	var name string
	var members string
	var creationDate string
	var firstAlbum string
	var location string
	var concertDates string
	var relations string
	for key, value := range tab.([]interface{})[0].(map[string]interface{}) {
		if key == "name" {
			name = fmt.Sprint(value)
		}
		if key == "members" {
			members = fmt.Sprint(value)
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
	}
	data := Page{name, members, creationDate, firstAlbum, location, concertDates, relations}
	return data
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	data := get(lien + "/artists")
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// affiche l'html
	tmpl, err := template.ParseFiles("./templates/index.gohtml")

	if err != nil {
	}
	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
