package building

type Parking struct {
	GarageType string
	PrkgSize   float32
	PrkgType   string
}

func NewParking(garageType string, prkgSize float32, prkgType string) Parking {
	return Parking{
		GarageType: garageType,
		PrkgSize:   prkgSize,
		PrkgType:   prkgType,
	}
}
