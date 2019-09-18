package building

//Rooms ...
type Rooms struct {
	BathFixtures int     `json:"bathFixtures"`
	Baths1qtr    int     `json:"baths1qtr"`
	Baths3qtr    int     `json:"baths3qtr"`
	BathsCalc    float32 `json:"bathsCalc"`
	BathsFull    int     `json:"bathsFull"`
	BathsHalf    int     `json:"bathsHalf"`
	BathsTotal   float32 `json:"bathsTotal"`
	Beds         int     `json:"beds"`
	RoomsTotal   int     `json:"roomsTotal"`
}

//NewRooms ...
func NewRooms(
	bathFixtures int,
	baths1qtr int,
	baths3qtr int,
	bathsCalc float32,
	bathsFull int,
	bathsHalf int,
	bathsTotal float32,
	beds int,
	roomsTotal int) Rooms {
	return Rooms{
		BathFixtures: bathFixtures,
		Baths1qtr:    baths1qtr,
		Baths3qtr:    baths3qtr,
		BathsCalc:    bathsCalc,
		BathsFull:    bathsFull,
		BathsHalf:    bathsHalf,
		BathsTotal:   bathsTotal,
		Beds:         beds,
		RoomsTotal:   roomsTotal,
	}
}
