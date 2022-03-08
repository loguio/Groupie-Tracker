package main

//on Importe toute les bibliothèques que l'on a besoin

type Page2 struct {
	Noyau    []ArtistAPI
	Page     int
	NbArtist int
}
type Page3 struct {
	Noyau    []TrieName
	Page     int
	NbArtist int
}
type Page4 struct {
	Noyau    []TrieDate
	Page     int
	NbArtist int
}
type Page5 struct {
	Noyau    []TrieMembers
	Page     int
	NbArtist int
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

type TrieDate struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Id           int    `json:"id"`
	CreationDate int    `json:"creationDate"`
	Page         int
}
type TrieMembers struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Id      int      `json:"id"`
	Members []string `json:"members"`
	Page    int
}
