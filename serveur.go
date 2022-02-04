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
	Title string
}

type Todo struct {
	Artistslink   string `json:"artists"`
	Locationslink string `json:"locations"`
	Dateslink     string `json:"dates"`
	Relationlink  string `json:"relation"`
}

type artist struct {
	Id           string
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

func get(adress string) {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var tab interface{}
	// Convert response body to Todo struct
	json.Unmarshal(bodyBytes, &tab)
	for key, value := range tab.([]interface{})[0].(map[string]interface{}) {
		fmt.Println(key, " :", value)
	}
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	get(lien + "/artists")
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// affiche l'html
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	data := Page{"PLZZZZZZZZZZZZ"}
	if err != nil {
	}
	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", data)
	})

	/////////////////////////////////

	fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur

}
