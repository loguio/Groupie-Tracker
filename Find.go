package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func rechercheFind(adress string, name []string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var tempo []string
	var artists []TrieName // nos artistes seront stockés dans cette variables
	var pageFind PageListArtist
	var oneArtist TrieName //on stock les données de un artiste danc cette variable
	for idArtist != 0 {    // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
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
		idArtist++
		tempo = strings.Split(strings.ToUpper(oneArtist.Name), "")
		for i := 0; i < len(name); i++ {
			if name[i] == tempo[i] {
			} else {
				break
			}
			if i == len(name)-1 {
				artists = append(artists, oneArtist)
			}
		}
		pageFind.Noyau = artists
	}
	var err error
	return pageFind, err
}
