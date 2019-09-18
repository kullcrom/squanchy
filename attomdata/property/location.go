package property

//Location ...
type Location struct {
	Accuracy  string  `json:"accuracy"`
	Elevation float32 `json:"elevation"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Distance  float32 `json:"distance"`
	GeoID     string  `json:"geoid"`
}

//NewLocation ...
func NewLocation(
	accuracy string,
	elevation float32,
	latitude string,
	longitude string,
	distance float32,
	geoID string) Location {
	return Location{
		Accuracy:  accuracy,
		Elevation: elevation,
		Latitude:  latitude,
		Longitude: longitude,
		Distance:  distance,
		GeoID:     geoID,
	}
}
