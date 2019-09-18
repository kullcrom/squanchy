package assessment

//Assessed ...
type Assessed struct {
	AssdImprValue float32 `json:"assdImprValue"`
	AssdLandValue float32 `json:"assdLandValue"`
	AssdTTLValue  float32 `json:"assdTtlValue"`
}

//NewAssessed ...
func NewAssessed(assdImprValue float32, assdLandValue float32, assdTTLValue float32) Assessed {
	return Assessed{
		AssdImprValue: assdImprValue,
		AssdLandValue: assdLandValue,
		AssdTTLValue:  assdTTLValue,
	}
}
