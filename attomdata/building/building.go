package building

//Building ...
type Building struct {
	Size         Size         `json:"size"`
	Rooms        Rooms        `json:"rooms"`
	Interior     Interior     `json:"interior"`
	Construction Construction `json:"construction"`
	Parking      Parking      `json:"parking"`
	Summary      Summary      `json:"summary"`
}

//NewBuilding ...
func NewBuilding(
	size Size,
	rooms Rooms,
	interior Interior,
	construction Construction,
	parking Parking,
	summary Summary) Building {
	return Building{
		Size:         size,
		Rooms:        rooms,
		Interior:     interior,
		Construction: construction,
		Parking:      parking,
		Summary:      summary,
	}
}
