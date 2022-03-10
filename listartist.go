package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func listartist(w http.ResponseWriter, r *http.Request) {
	lien := "https://groupietrackers.herokuapp.com/api"
	page := 1
	var data interface{}
	nbArtist := 12
	tmpl, err := template.ParseFiles("./templates/navbar.html", "./templates/footer.html", "./templates/pagelistartists.html", "./templates/listartist.html") // utilisation du fichier navPage.gohtml pour le template
	if err != nil {
		error500(err, w)
	} else {
		if r.FormValue("FilterAlpha") == "TRUE" { // récupération de la variable qui nous indique quelle filtre on utilise
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				function := "FilterAlpha"
				data, err = FilterAlpha(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
				if err != nil {
					error500(err, w)
				}
			}
		} else if r.FormValue("FilterDate") == "TRUE" { // récupération de la variable qui nous indique quelle filtre on utilise
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				function := "FilterDate"
				data, err = FilterDate(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
				if err != nil {
					error500(err, w)
				}
			}
		} else if r.FormValue("FilterSolo") == "TRUE" { // récupération de la variable qui nous indique quelle filtre on utilise
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				function := "FilterSolo"
				data, err = FilterSolo(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
				if err != nil {
					error500(err, w)
				}
			}
		} else if r.FormValue("FilterGroups") == "TRUE" { // récupération de la variable qui nous indique quelle filtre on utilise
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				function := "FilterGroups"
				data, err = FilterGroups(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
				if err != nil {
					error500(err, w)
				}
			}
		} else if r.FormValue("pageSuivante") == "TRUE" { // avancé a la page suivante en gardant les filtre si utilisé grâce a la variable function
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				page, err = strconv.Atoi(r.FormValue("page"))
				if err != nil {
					error500(err, w)
				} else {
					page += 1
					if r.FormValue("function") == "FilterSolo" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterSolo(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterGroups" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterGroups(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterAlpha" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterAlpha(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterDate" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterDate(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							fmt.Println(err)
						}
					} else {
						function := "normal"
						data, err = ArtistPage(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					}
				}
			}
		} else if r.FormValue("pagePrecedente") == "TRUE" { // avancé a la page précédente en gardant les filtre si utilisé grâce a la variable function
			nbArtist, err = strconv.Atoi(r.FormValue("Artists")) // récupération du nombre d'artists à afficher
			if err != nil {
				error500(err, w)
			} else {
				page, err = strconv.Atoi(r.FormValue("page"))
				if err != nil {
					error500(err, w)
				} else {
					page -= 1
					if r.FormValue("function") == "FilterSolo" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterSolo(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterGroups" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterGroups(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterAlpha" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterAlpha(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					} else if r.FormValue("function") == "FilterDate" { // récupération de la variable qui nous indique quelle filtre on utilise
						function := r.FormValue("function")
						data, err = FilterDate(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							fmt.Println(err)
						}
					} else {
						function := "normal"
						data, err = ArtistPage(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
						if err != nil {
							error500(err, w)
						}
					}
				}
			}
		} else {
			function := "normal"
			data, err = ArtistPage(lien+"/artists", page, nbArtist, function) //  récupération des artists en fonction du filtre
			if err != nil {
				error500(err, w)
			}
		}
		tmpl.ExecuteTemplate(w, "listartists", data) //exécution du template
	}
}
