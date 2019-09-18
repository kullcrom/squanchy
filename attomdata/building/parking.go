package building

//Parking ...
type Parking struct {
	GarageType string  `json:"garageType"`
	PrkgSize   float32 `json:"prkgSize"`
	PrkgType   string  `json:"prkgType"`
}

//NewParking ...
func NewParking(garageType string, prkgSize float32, prkgType string) Parking {
	return Parking{
		GarageType: garageType,
		PrkgSize:   prkgSize,
		PrkgType:   prkgType,
	}
}
