package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func lieux(address string) []string {
	var Locate Location
	Locate.Id = 1
	var listLieux []string
	for Locate.Id != 0 {
		var tempo string
		tempo = strconv.Itoa(Locate.Id)
		resp, err := http.Get(address + tempo) // on recupère les données qui sont stockés dans resp
		if err != nil {
			log.Fatalln(err) // si il y a une erreur donc erreur
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, &Locate)
		for i := 0; i < len(Locate.Location); i++ {
			listLieux = append(listLieux, bonLieu(Locate.Location[i]))
		}
		if Locate.Id == 0 {
			break
		}
		Locate.Id++
	}
	return listLieux
}
