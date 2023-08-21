package pkg

import (
	"strconv"
	"strings"
	"zsandibe/models"
)

var (
	ArtistInterface    = []models.Artist{}
	RelationsInterface = models.Relation{}
)

func SearchGroup(data string) (*models.Result, error) {
	var result []models.Artist
	artists, err := GiveData(&ArtistInterface, &RelationsInterface)
	// fmt.Println(err)
	if err != nil {
		return &models.Result{}, err
	}
	for _, artist := range artists.Artists {
		var contains bool
		if MyContains(artist.Name, data) {
			result = append(result, artist)
			continue
		} else if strings.Contains(strconv.Itoa(artist.CreationDate), data) {
			result = append(result, artist)
			continue
		} else if strings.Contains(artist.FirstAlbum, data) {
			result = append(result, artist)
			continue
		} else {
			for _, m := range artist.Members {
				if strings.Contains(m, data) {
					result = append(result, artist)
					contains = true
					break
				}
			}
			if contains {
				continue
			}
		}

		for city, dates := range artist.DatesLocations {
			if MyContains(city, data) {
				result = append(result, artist)
				break
			}
			for _, date := range dates {
				if MyContains(date, data) {
					result = append(result, artist)
					contains = true
					break
				}
			}
			if contains {
				break
			}
		}
	}
	ResultSearching := models.Result{Artists: artists.Artists, FoundArtists: result}
	// fmt.Println(ResultSearching)

	return &ResultSearching, nil
}

// function for search in anycase
func MyContains(s string, substr string) bool {
	// fmt.Println(s, " :s ", substr, " :substr ", " ", strings.Contains(strings.ToLower(s), strings.ToLower(substr)))
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func IsPrintable(text string) bool {
	printableFlag := false
	for _, char := range text {
		if char != ' ' && char != '\t' && char != '\n' && char != '\r' {
			printableFlag = true
			if char < 32 || char > 126 {
				return false
			}
		}
	}
	return printableFlag
}
