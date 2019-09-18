package property

//Area ...
type Area struct {
	CountrySecSubd   string `json:"countrySecSubd"`
	SubdName         string `json:"subdName"`
	SubdTractNum     string `json:"subdTractNum"`
	CensusTractIdent string `json:"censusTractIdent"`
	CensusBlockGroup string `json:"censusBlockGroup"`
}

//NewArea ...
func NewArea(countrySecSubd string, subdName string, subdTractNum string, censusTractIdent string, censusBlockGroup string) Area {
	return Area{
		CountrySecSubd:   countrySecSubd,
		SubdName:         subdName,
		SubdTractNum:     subdTractNum,
		CensusTractIdent: censusTractIdent,
		CensusBlockGroup: censusBlockGroup,
	}
}
