package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
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
	Relations           []string
}

type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	Id             int `json:"id"`
	DatesLocations struct {
	}
}

func HomePage(adress string, nbPage int) (interface{}, error) {
	var idArtist = (nbPage-1)*12 + 1
	var url = ""
	var artists []ArtistAPI
	var oneartist ArtistAPI
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	for idArtist != nbPage*12+1 {
		url = "/" + strconv.Itoa(idArtist)
		resp, err := http.Get(adress + url)
		if err != nil {
			fmt.Println(err)
			return artists, err
		} else {
			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			json.Unmarshal(bodyBytes, &oneartist)
			idArtist = oneartist.Id
			if idArtist == 0 {
				fmt.Println("Error : API vide")
				err = errors.New("Invalid API Id")
				return artists, err
			}
			oneartist.Location, err = location(oneartist.AddressLocation)
			if err != nil {
				return artists, err
			}
			oneartist.ConcertDates, err = concertdate(oneartist.ConcertDatesaddress)
			if err != nil {
				return artists, err
			}
			artists = append(artists, oneartist)
			idArtist++
			//fmt.Println(oneartist)
		}
	}
	var err error
	return artists, err
}
func concertdate(adress string) ([]string, error) {
	var dates Dates
	var mois []string
	var date string
	resp, err := http.Get(adress)
	if err != nil {
		fmt.Println(err)
		return mois, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &dates)
	for i := 0; i < len(dates.Dates); i++ {
		mois = strings.Split(dates.Dates[i], "-")
		if mois[1] == "01" {
			date = strings.Replace(mois[1], "01", "Janvier", -1)
		} else if mois[1] == "02" {
			date = strings.Replace(mois[1], "02", "Fevrier", -1)
		} else if mois[1] == "03" {
			date = strings.Replace(mois[1], "03", "Mars", -1)
		} else if mois[1] == "04" {
			date = strings.Replace(mois[1], "04", "Avril", -1)
		} else if mois[1] == "05" {
			date = strings.Replace(mois[1], "05", "Mai", -1)
		} else if mois[1] == "06" {
			date = strings.Replace(mois[1], "06", "Juin", -1)
		} else if mois[1] == "07" {
			date = strings.Replace(mois[1], "07", "Juillet", -1)
		} else if mois[1] == "08" {
			date = strings.Replace(mois[1], "08", "Aout", -1)
		} else if mois[1] == "09" {
			date = strings.Replace(mois[1], "09", "Septembre", -1)
		} else if mois[1] == "10" {
			date = strings.Replace(mois[1], "10", "Octobre", -1)
		} else if mois[1] == "11" {
			date = strings.Replace(mois[1], "11", "Novembre", -1)
		} else if mois[1] == "12" {
			date = strings.Replace(mois[1], "12", "Decembre", -1)
		}
		mois[1] = date
		date = strings.Join(mois, " ")
		date = strings.Replace(date, "*", "", -1)
		dates.Dates[i] = date
	}
	return dates.Dates, err
}

func location(adress string) ([]string, error) {
	var locations Location
	var location string
	resp, err := http.Get(adress)
	if err != nil {
		fmt.Println(err)
		return locations.Location, err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &locations)
	for i := 0; i < len(locations.Location); i++ {
		location = locations.Location[i]
		location = strings.Replace(location, "_", " ", -1)
		location = strings.Replace(location, "-", " ", -1)
		locations.Location[i] = location
	}
	return locations.Location, err
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func main() {
	lien := "https://groupietrackers.herokuapp.com/api"
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/assets/", http.StripPrefix("/assets/", fileServer))
	// affiche l'html
	page := 1
	http.HandleFunc("/Groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/Groupie-tracker" {
			fmt.Fprintln(w, "uwu")
			errorHandler(w, r, http.StatusNotFound)
		} else {
			data, err := HomePage(lien+"/artists", page)
			if err != nil {
				tmpl, err := template.ParseFiles("./Error500.gohtml")
				if err != nil {
				}
				tmpl.ExecuteTemplate(w, "index", data)
			} else {
				tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
				if err != nil {
					tmpl, err = template.ParseFiles("./Error500.gohtml")
					if err != nil {
					}
					tmpl.ExecuteTemplate(w, "index", data)
				} else {
					tmpl.ExecuteTemplate(w, "index", data)
				}
			}
		}

	})

	http.HandleFunc("/Groupie-tracker/PageSuivante", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/Groupie-tracker/PageSuivante" {
			fmt.Fprintln(w, "uwusgsg")
			errorHandler(w, r, http.StatusNotFound)
		} else {
			if r.Method == "POST" {
				page += 1
				data, err := HomePage(lien+"/artists", page)
				if err != nil {
					tmpl, err := template.ParseFiles("./Error500.gohtml")
					if err != nil {
					}
					tmpl.ExecuteTemplate(w, "index", data)
				} else {
					tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
					if err != nil {
						tmpl, err = template.ParseFiles("./Error500.gohtml")
						if err != nil {
						}
						tmpl.ExecuteTemplate(w, "index", data)
					} else {
						tmpl.ExecuteTemplate(w, "index", data)
					}
				}
			}
		}
	})

	http.HandleFunc("/Groupie-tracker/PagePrecedente", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
		}
		page -= 1
		data, err := HomePage(lien+"/artists", page)
		if err != nil {
			tmpl, err := template.ParseFiles("./Error500.gohtml")
			if err != nil {
			}
			tmpl.ExecuteTemplate(w, "index", data)
		} else {
			tmpl, err := template.ParseFiles("./assets/navPage.gohtml")
			if err != nil {
				tmpl, err = template.ParseFiles("./Error500.gohtml")
				if err != nil {
				}
				tmpl.ExecuteTemplate(w, "index", data)
			} else {
				tmpl.ExecuteTemplate(w, "index", data)
			}
		}
	})

	fmt.Println("le serveur est en cours d'éxécution a l'adresse http://localhost:3000/Groupie-tracker")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur
}
