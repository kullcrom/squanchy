package property

type Utilities struct {
	CoolingType string
	EnergyType  string
	WallType    string
}

func NewUtilities(coolingType string, energyType string, wallType string) Utilities {
	return Utilities{
		CoolingType: coolingType,
		EnergyType:  energyType,
		WallType:    wallType,
	}
}
