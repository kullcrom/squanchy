package building

type Rooms struct {
	BathFixtures int
	Baths1qtr    int
	Baths3qtr    int
	BathsCalc    float32
	BathsFull    int
	BathsHalf    int
	BathsTotal   float32
	Beds         int
	RoomsTotal   int
}

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
