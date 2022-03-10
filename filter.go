package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func FilterAlpha(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page PageListArtist
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
	sort.Strings(names) //trie les noms par ordre alphabétique
	for i := (Page-1)*nbArtist + 1; i != len(names); i++ {
		for j := 1; j != len(names); j++ {
			if names[i] == tempartists[j].Name {
				artists = append(artists, tempartists[j])
				break
			}
		} // double boucle pour que artist sois trier
		if len(artists) == nbArtist {
			break
		} // arreter la boucle quand artiste a atteint le nombre d'artist à afficher sur la page
	}
	var err error
	page.Noyau = artists
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}

func FilterDate(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	var page PageFilterDate
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

	}
	sort.Ints(creationDate) //trie les noms par ordre de date de création
	for i := 0; i != len(creationDate); i++ {
		for j := 0; j != len(tempartists); j++ {
			if creationDate[i] == tempartists[j].CreationDate {
				artists = append(artists, tempartists[j])
				tempartists = remove(tempartists, j)
				break
			}
		}
	} // double boucle qui trie artist dans l'ordre de la date de création
	var tempo2 []TrieDate
	tempo := nbArtist * (Page - 1)
	for i := tempo; i < nbArtist*Page; i++ {
		tempo2 = append(tempo2, artists[i])
	} //renvoie le nombre d'artiste demander
	var err error
	page.Noyau = tempo2
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}

func FilterGroups(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	point := 0
	var page PageFilterMembers
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
			if point > (Page-1)*nbArtist { // prendre seulement les artists qu'on affiche sur la page
				artists = append(artists, oneArtist)
			} else {
				point++
			}
		} // prendre que les groups
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

func FilterSolo(adress string, Page int, nbArtist int, function string) (interface{}, error) {
	var idArtist = 1 // on prend le première identifiant de l'artiste que l'utilisateur veut afficher
	var url = ""
	point := 1
	var page PageFilterMembers
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
		if len(oneArtist.Members) == 1 { // prendre que les artists solo
			if point > (Page-1)*nbArtist { // commencer a prendre que les artist qu'on affiche sur la page a la bonne page
				artists = append(artists, oneArtist)
			} else {
				point++
			}
		}
		idArtist++
		if len(artists) == nbArtist {
			break
		} // prend le nombre d'artists demandé
	}
	var err error
	page.Noyau = artists
	page.Function = function
	page.Page = Page
	page.NbArtist = nbArtist
	return page, err
}
