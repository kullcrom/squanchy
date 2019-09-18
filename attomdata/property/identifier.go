package property

//Identifier ...
type Identifier struct {
	ObPropID uint64 `json:"obPropId"`
	Fips     string `json:"fips"`
	Apn      string `json:"apn"`
	AttomID  uint64 `json:"attomId"`
}

//NewIdentifier ...
func NewIdentifier(obPropID uint64, fips string, apn string, attomID uint64) Identifier {
	return Identifier{
		ObPropID: obPropID,
		Fips:     fips,
		Apn:      apn,
		AttomID:  attomID,
	}
}
