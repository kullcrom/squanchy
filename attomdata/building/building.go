package building

type Building struct {
	size         Size
	rooms        Rooms
	interior     Interior
	construction Construction
	parking      Parking
	summary      Summary
}

func getSize() Size {
	return this.Size
}
