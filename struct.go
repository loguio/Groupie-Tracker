package main

type PageListArtist struct {
	Noyau    []TrieName
	Page     int
	NbArtist int
	Function string
}
type PageFilterName struct { //Structure des donnée envoyer sur l'html
	Noyau    []TrieName
	Page     int
	NbArtist int
	Function string
}
type PageFilterDate struct { //Structure des donnée envoyer sur l'html
	Noyau    []TrieDate
	Page     int
	NbArtist int
	Function string
}
type PageFilterMembers struct { //Structure des donnée envoyer sur l'html
	Noyau    []TrieMembers
	Page     int
	NbArtist int
	Function string
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
	Relations           map[string][]string
	RelationDate        [][]string
	DateLocation        []DateLocation
} // on créer une structure qui contient toutes les données pouvant etre utile a notre site cela va nous permettre d'afficher chaque groupe avec leur données respectives

type DateLocation struct {
	Location string
	Dates    []string
} //Cette structure nous permet d'afficher les lieu ainsi que les dates des spectacles

type Carte struct {
	Valeur   string
	Location []string
}

type Location struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Location

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
} // cette strcture nous permet de recuperer les donnée du lien API Dates

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
} // cette strcture nous permet de recuperer les donnée du lien API Relation

type TrieName struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Id    int    `json:"id"`
	Page  int
}

type TrieDate struct { // cette structure récupere les donnée utile pour le filtre de date
	Name         string `json:"name"`
	Image        string `json:"image"`
	Id           int    `json:"id"`
	CreationDate int    `json:"creationDate"`
	Page         int
}
type TrieMembers struct { // cette structure récupere les donnée utile pour le filtre de membres
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Id      int      `json:"id"`
	Members []string `json:"members"`
	Page    int
}
type Page6 struct {
	Noyau []TrieName
}
