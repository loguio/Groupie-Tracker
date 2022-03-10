package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func trieAlpha(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page Page3
	var names []string
	var tempartists []TrieName // nos artistes seront stockés dans cette variables
	var artists []TrieName     // nos artistes seront stockés dans cette variables
	var oneArtist TrieName     //on stock les données de un artiste danc cette variable
	for idArtist != 0 {        // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
		url = "/" + strconv.Itoa(idArtist)  // On recupère un URL equivalent a afficher les données avec l'ID d'un artiste
		resp, err := http.Get(adress + url) // on recupère les données qui sont stockés dans resp
		if err != nil {
			fmt.Println(err) // si il y a une erreur donc erreur
			return tempartists, err
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
		idArtist = oneArtist.Id
		if idArtist == 0 {
			fmt.Println("erreur : L'API est vide")
			break
		} // si l'id est égal a 0 c'est que l'on a atteint la fin des artistes et que il n'y en a pas plus a afficher donc on return pour sortir de la boucle
		tempartists = append(tempartists, oneArtist)
		idArtist++
		names = append(names, oneArtist.Name)
	}
	sort.Strings(names)
	for i := (Page-1)*nbArtist + 1; i != len(names); i++ {
		for j := 1; j != len(names); j++ {
			if names[i] == tempartists[j].Name {
				artists = append(artists, tempartists[j])
				break
			}
		}
		if len(artists) == nbArtist {
			break
		}
	}
	var err error
	page.Noyau = artists
	page.Function = function
	fmt.Println(page.Function)
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}

func trieDate(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = (Page-1)*nbArtist + 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page Page4
	var creationDate []int
	var tempartists []TrieDate // nos artistes seront stockés dans cette variables
	var artists []TrieDate     // nos artistes seront stockés dans cette variables
	var oneArtist TrieDate     //on stock les données de un artiste danc cette variable
	for idArtist != 0 {        // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
		url = "/" + strconv.Itoa(idArtist)  // On recupère un URL equivalent a afficher les données avec l'ID d'un artiste
		resp, err := http.Get(adress + url) // on recupère les données qui sont stockés dans resp
		if err != nil {
			fmt.Println(err) // si il y a une erreur donc erreur
			return tempartists, err
		}
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &oneArtist) // on implémente les données contenus dans bodyBytes dans oneArtist cela va nous permettre de recupérer les données
		idArtist = oneArtist.Id
		if idArtist == 0 {
			fmt.Println("erreur : L'API est vide")
			break
		} // si l'id est égal a 0 c'est que l'on a atteint la fin des artistes et que il n'y en a pas plus a afficher donc on return pour sortir de la boucle
		tempartists = append(tempartists, oneArtist)
		idArtist++
		creationDate = append(creationDate, oneArtist.CreationDate)
		if len(tempartists) == nbArtist {
			break
		}
	}
	sort.Ints(creationDate)
	for i := 0; i != len(creationDate); i++ {
		for j := 0; j != len(tempartists); j++ {
			if creationDate[i] == tempartists[j].CreationDate {
				artists = append(artists, tempartists[j])
				tempartists = remove(tempartists, j)
				break
			}
		}
	}
	var err error
	page.Noyau = artists
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}

func trieGroups(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = (Page-1)*nbArtist + 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page Page5
	var artists []TrieMembers // nos artistes seront stockés dans cette variables
	var oneArtist TrieMembers //on stock les données de un artiste danc cette variable
	for idArtist != 0 {       // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
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
		if len(oneArtist.Members) > 1 {
			artists = append(artists, oneArtist)
		}
		idArtist++
		if len(artists) == nbArtist {
			break
		}
	}
	var err error
	page.Noyau = artists
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}

func trieSolo(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = (Page-1)*nbArtist + 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page Page5
	var artists []TrieMembers // nos artistes seront stockés dans cette variables
	var oneArtist TrieMembers //on stock les données de un artiste danc cette variable
	for idArtist != 0 {       // on repete cette action jusqu'a ce qu'on ait recupéré les données de 12 artistes
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
		if len(oneArtist.Members) == 1 {
			artists = append(artists, oneArtist)
		}
		idArtist++
		if len(artists) == nbArtist {
			break
		}
	}
	var err error
	page.Noyau = artists
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}
