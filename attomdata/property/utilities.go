package property

//Utilities ...
type Utilities struct {
	CoolingType string `json:"coolingType"`
	EnergyType  string `json:"energyType"`
	WallType    string `json:"wallType"`
}

//NewUtilities ...
func NewUtilities(coolingType string, energyType string, wallType string) Utilities {
	return Utilities{
		CoolingType: coolingType,
		EnergyType:  energyType,
		WallType:    wallType,
	}
}
