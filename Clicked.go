package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func clicked(id string) (interface{}, error) {
	var oneArtist ArtistAPI
	var relationdate [][]string
	var clean []DateLocation
	url := "https://groupietrackers.herokuapp.com/api/artists/" + id
	resp, err := http.Get(url) // on recupère les données qui sont stockés dans resp
	if err != nil {
		log.Fatalln(err) // si il y a une erreur donc erreur
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &oneArtist)                         // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
	oneArtist.FirstAlbum = gooddate(oneArtist.FirstAlbum)         //on passe la donnée FirstAlbum dans la fonction gooddate pour avoir une date plus explicite
	oneArtist.Location, err = location(oneArtist.AddressLocation) // on Récupere les données qui nous interesse grace a la fonction Location car AddressLocation est un lien API
	if err != nil {
		return oneArtist, err
	}
	oneArtist.ConcertDates, err = concertdate(oneArtist.ConcertDatesaddress) // ConcertDatesaddress est aussi un lien API donc on tri les données pur avoir celle qui nous intéresse
	if err != nil {
		return oneArtist, err
	}
	oneArtist.Relations, err = relation(oneArtist.RelationsAdress) // on fait la meme chose pour RelationsAdress
	if err != nil {
		return oneArtist, err
	}
	for i := 0; i < len(oneArtist.Location); i++ {
		for k := 0; k < len(oneArtist.Relations[oneArtist.Location[i]]); k++ {
			oneArtist.Relations[oneArtist.Location[i]][k] = gooddate(oneArtist.Relations[oneArtist.Location[i]][k]) //on change aussi la date des concert our avoir une date plus joli
		}
		relationdate = append(relationdate, oneArtist.Relations[oneArtist.Location[i]]) // on rajoute les valeurs des dates dans l'index de la villes correspondante
	}
	oneArtist.RelationDate = relationdate // on stock les valeurs des dates dans
	oneArtist.DateLocation = clean        // on vide notre liste
	for i := 0; i < len(oneArtist.Location); i++ {
		var tempo DateLocation
		tempo.Location = bonLieu(oneArtist.Location[i])
		tempo.Dates = oneArtist.RelationDate[i]
		oneArtist.DateLocation = append(oneArtist.DateLocation, tempo) // on ajoute les valeurs utile dans DateLocation
	}
	return oneArtist, err
}
