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
	Relations	map[string][]string
	RelationDate [][]string
	Test []Machin
}
type Machin struct {
	Location string
	Dates []string
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
	DatesLocations map[string][]string `json:"datesLocations"`
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
		var relationdate [][]string
		url = "/"+strconv.Itoa(idArtist)
		resp, err := http.Get(adress+url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneartist)
		idArtist = oneartist.Id 
		if idArtist == 0 {break}
		oneartist.FirstAlbum = gooddate(oneartist.FirstAlbum)
		oneartist.Location = location(oneartist.AddressLocation)
		oneartist.ConcertDates = concertdate(oneartist.ConcertDatesaddress)
		oneartist.Relations = relation(oneartist.RelationsAdress)
		for i:=0;i<len(oneartist.Location);i++ {
			for k := 0; k< len(oneartist.Relations[oneartist.Location[i]]);k++ {
				oneartist.Relations[oneartist.Location[i]][k]  = gooddate(oneartist.Relations[oneartist.Location[i]][k])
			}
			relationdate=append(relationdate,oneartist.Relations[oneartist.Location[i]])
		}
		oneartist.RelationDate = relationdate
		for i:=0; i < len(oneartist.Location);i++ {
			var Test50 Machin
			Test50.Location = oneartist.Location[i]
			Test50.Dates = oneartist.RelationDate[i]
			oneartist.Test = append(oneartist.Test,Test50)
		}
		artists = append(artists,oneartist)
		fmt.Println(oneartist)
		idArtist++
	}
	return artists
}

func relation(adress string) map[string][]string {
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
