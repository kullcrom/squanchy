package property

type Area struct {
	CountrySecSubd   string
	SubdName         string
	SubdTractNum     string
	CensusTractIdent string
	CensusBlockGroup string
}

func NewArea(countrySecSubd string, subdName string, subdTractNum string, censusTractIdent string, censusBlockGroup string) Area {
	return Area{
		CountrySecSubd:   countrySecSubd,
		SubdName:         subdName,
		SubdTractNum:     subdTractNum,
		CensusTractIdent: censusTractIdent,
		CensusBlockGroup: censusBlockGroup,
	}
}
