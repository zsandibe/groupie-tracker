package models

type Relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"dateLocations"`
	}
}
