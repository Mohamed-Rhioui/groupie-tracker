package roots

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

// Define the structure that matches the JSON response
type Artist struct {
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	Image           string       `json:"image"`
	CreationDate    int          `json:"creationDate"`
	Members         []string     `json:"members"`
	FirstAlbum      string       `json:"firstAlbum"`
	LocationsURL    string       `json:"locations"`
	ConcertDatesURL string       `json:"concertDates"`
	RelationsURL    string       `json:"relations"`
	Locations       Locations    `json:"locationsData"`
	ConcertDates    ConcertDates `json:"concertDatesData"`
	Relations       Relations    `json:"relationsData"`
}

type Locations struct {
	Locations []string `json:"locations"`
}

type ConcertDates struct {
	Dates []string `json:"dates"`
}

type Relations struct {
	RelatedArtists map[string][]string `json:"datesLocations"`
}

func HandleMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	ArtistsUrl := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(ArtistsUrl)
	if err != nil {
		fmt.Printf("The HTTP request failed with error: %s\n", err)
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		return
	}

	var artists []Artist
	err = json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %s\n", err)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleDetailsPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/details.html"))

	LocationsURL := "https://groupietrackers.herokuapp.com/api/locations"
	dates := "https://groupietrackers.herokuapp.com/api/dates"
	relation := "https://groupietrackers.herokuapp.com/api/relation"
}
