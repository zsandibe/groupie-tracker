package pkg

import (
	"strconv"
	"zsandibe/models"
)

func FilterGroup(minCreation, maxCreation, minAlbum, maxAlbum, location string, members []string) (*models.Result, error) {
	var foundArtist []models.Artist
	// fmt.Println(minCreation, maxCreation, minAlbum, maxAlbum, location, members)
	minCreation_int, _ := strconv.Atoi(minCreation)
	maxCreation_int, _ := strconv.Atoi(maxCreation)
	minAlbum_int, _ := strconv.Atoi(minAlbum)
	maxAlbum_int, _ := strconv.Atoi(maxAlbum)
	artists, _ := GiveData(&ArtistInterface, &RelationsInterface)

	for _, a := range artists.Artists {
		album, _ := strconv.Atoi(a.FirstAlbum[6:])
		var flag1, flag2, flag3 bool
		if (a.CreationDate <= maxCreation_int && a.CreationDate >= minCreation_int) && (album <= maxAlbum_int && album >= minAlbum_int) {
			flag1 = true
		}
		if len(location) == 0 {
			flag2 = true
		}

		for city := range a.DatesLocations {
			// city = strings.Replace(strings.Replace(city, "_", " ", -1), "-", ", ", -1)
			// fmt.Println(city, location)
			if MyContains(city, location) {
				flag2 = true
			}
		}
		for _, w := range members {
			s, _ := strconv.Atoi(w)
			if len(a.Members) == s {
				flag3 = true
			}
		}
		if len(members) == 0 {
			flag3 = true
		}
		if flag1 && flag2 && flag3 {
			foundArtist = append(foundArtist, a)
		}

	}
	result := &models.Result{Artists: artists.Artists, FoundArtists: foundArtist}
	return result, nil
}
