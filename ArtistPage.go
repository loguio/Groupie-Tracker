package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ArtistPage(adress string, Page int, nbArtist int) (interface{}, error) { //Cette fonction se lance lorsque l'utilisateur est sur la page des artistes
	fmt.Println("1. Performing Http Get...")
	var idArtist = (Page-1)*nbArtist + 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page Page2
	var artists []ArtistAPI // nos artistes seront stockés dans cette variables
	var oneArtist ArtistAPI //on stock les données de un artiste danc cette variable
	fmt.Println("1. Performing Http Get...")
	fmt.Println("2. Le serveur est lancé sur le port 3000")
	for idArtist != Page*nbArtist+1 { // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
		url = "/" + strconv.Itoa(idArtist)  // On recupère un URL equivalent a afficher les données avec l'ID d'un artiste
		resp, err := http.Get(adress + url) // on recupère les données qui sont stockés dans resp
		if err != nil {
			fmt.Println(err) // si il y a une erreur donc erreur
			return artists, err
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
		idArtist = oneArtist.Id
		if idArtist == 0 {
			fmt.Println("erreur : L'API est vide")
			break
		} // si l'id est égal a 0 c'est que l'on a atteint la fin des artistes et que il n'y en a pas plus a afficher donc on return pour sortir de la boucle
		if err != nil {
			return artists, err
		}
		artists = append(artists, oneArtist)
		idArtist++
	}
	var err error
	page.Noyau = artists
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err // on renvois notre liste avec 12 artistes et les données
}
