package property

type Location struct {
	Accuracy  string
	Elevation float32
	Latitude  string
	Longitude string
	Distance  float32
	GeoID     string
}

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
