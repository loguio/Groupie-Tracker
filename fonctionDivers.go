package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func relation(adress string) (map[string][]string, error) {
	//relation nous permet de prendre les valeurs que l'on a vraimen;$t beosin dans L'API Relation
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
func remove(s []TrieDate, i int) []TrieDate {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
func bonLieu(date string) string {
	// cette fonctions permet d'enlever les caractères qui ne servent a rien
	date = strings.Replace(date, "-", " ", -1)
	date = strings.Replace(date, "_", " ", -1)
	return date
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
