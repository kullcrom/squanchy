package property

type Identifier struct {
	ObPropID uint64
	Fips     string
	Apn      string
	AttomID  uint64
}

func NewIdentifier(obPropID uint64, fips string, apn string, attomID uint64) Identifier {
	return Identifier{
		ObPropID: obPropID,
		Fips:     fips,
		Apn:      apn,
		AttomID:  attomID,
	}
}
