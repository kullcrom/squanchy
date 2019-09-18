package building

//Interior ...
type Interior struct {
	BsmtSize  int    `json:"bsmtSize"`
	FplcCount int    `json:"fplcCount"`
	FplcInd   string `json:"fplcInd"`
	FplcType  string `json:"fplcType"`
}

//NewInterior ...
func NewInterior(bsmtSize int, fplcCount int, fplcInd string, fplcType string) Interior {
	return Interior{
		BsmtSize:  bsmtSize,
		FplcCount: fplcCount,
		FplcInd:   fplcInd,
		FplcType:  fplcType,
	}
}
