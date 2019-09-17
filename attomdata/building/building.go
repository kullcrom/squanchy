package building

type Building struct {
	Size         Size
	Rooms        Rooms
	Interior     Interior
	Construction Construction
	Parking      Parking
	Summary      Summary
}

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
