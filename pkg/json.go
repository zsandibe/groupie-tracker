package pkg

import (
	"encoding/json"
	"net/http"
	"zsandibe/models"
)

var (
	artists  models.Result
	relation models.Relation
)

const (
	urlArtist   = "https://groupietrackers.herokuapp.com/api/artists"
	urlRelation = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetBody(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

func GiveData(artist *[]models.Artist, relation *models.Relation) (*models.Result, error) {
	if err := GetBody(urlArtist, &artists.Artists); err != nil {
		return &artists, err
	}

	if err := GetBody(urlRelation, &relation); err != nil {
		return &artists, err
	}

	for i, v := range relation.Index {
		artists.Artists[i].DatesLocations = v.DatesLocations
	}
	return &artists, nil
}
