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
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

func get(adress string) {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get(adress)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	get(lien)
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
