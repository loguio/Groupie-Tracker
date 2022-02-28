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
	Image        string
	Members      []string
	CreationDate string
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}
type Page2 struct {
	Name  []string
	Image []string
}

type Todo struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func HomePage(adress string) interface{} {
	fmt.Println("1. Performing Http Get...")
	var id = 1
	var test = ""
	var name []string
	var image []string
	var tab Todo
	for id != 0 {
		test = "/" + strconv.Itoa(id)
		resp, err := http.Get(adress + test)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &tab)
		name = append(name, tab.Name)
		image = append(image, tab.Image)
		id = tab.Id
		if id == 0 {
			break
		}
		id++
	}
	data := Page2{name, image}
	fmt.Println(data)
	return data
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
	data := Page{name, image, members, creationDate, firstAlbum, location, concertDates, relations}
	return data, valreturn
}

func main() {
	var fs = http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	lien := "https://groupietrackers.herokuapp.com/api"
	// fileServer := http.FileServer(http.Dir("assets/")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	// http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	// affiche l'html
	tmpl, err := template.ParseFiles("./assets/layout.html", "./assets/header.html", "./assets/footer.html")
	if err != nil {
	}
	nb := 0
	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/layout.html")
			if err != nil {
			}
			nb, _ = strconv.Atoi(r.FormValue("nombre"))
		}
		data, codeError := get(lien+"/artists", nb)
		if codeError == 500 {

		}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/artist", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/index.gohtml")
			nb, _ = strconv.Atoi(r.FormValue("nombre"))
		}
		data, codeError := get(lien+"/artists", nb)
		if codeError == 500 {
		}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	http.HandleFunc("/Groupie-tracker/artist=", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl, err = template.ParseFiles("./assets/artistes.gohtml")
			nb, _ = strconv.Atoi(r.FormValue("nombre"))
		}
		data, codeError := get(lien+"/artists", nb)
		if codeError == 500 {
		}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
