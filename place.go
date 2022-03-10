package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func place(address string) []string {
	// cette fonction permet de recupere tout les lieux de tout les artistes et de les renvoyer
	var Locate Location
	Locate.Id = 1
	var listPlace []string
	for Locate.Id != 0 {
		var id string
		id = strconv.Itoa(Locate.Id)
		resp, err := http.Get(address + id) // on recupère les données qui sont stockés dans resp
		if err != nil {
			log.Fatalln(err) // si il y a une erreur donc erreur
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &Locate)
		for i := 0; i < len(Locate.Location); i++ {
			listPlace = append(listPlace, bonLieu(Locate.Location[i])) // on ajoute toutes les valeurs de chaque artistes dans la variable listPlace
		}
		if Locate.Id == 0 { // si on atteint le bout des artistes on sort
			break // on sort de la boucle
		}
		Locate.Id++                   // on ajoute 1 pour changer d'artiste (passé au suivant)
		sort.Strings(listPlace)       // on trie la list avec la bibliothèque sort
		listPlace = double(listPlace) // onenlève les doublons
	}
	return listPlace
}

func double(list []string) []string {
	// cette fonction permet d'enlever les doublons d'une liste
	counts := make(map[string]bool)
	for _, x := range list {
		counts[x] = true
	}
	result := make([]string, len(counts))
	j := len(result) - 1
	for i := len(list) - 1; i >= 0; i-- {
		if counts[list[i]] {
			counts[list[i]] = false
			result[j] = list[i]
			j--
		}
	}
	return result
}
