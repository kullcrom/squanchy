package building

type Interior struct {
	BsmtSize  int
	FplcCount int
	FplcInd   string
	FplcType  string
}

func NewInterior(bsmtSize int, fplcCount int, fplcInd string, fplcType string) Interior {
	return Interior{
		BsmtSize:  bsmtSize,
		FplcCount: fplcCount,
		FplcInd:   fplcInd,
		FplcType:  fplcType,
	}
}
