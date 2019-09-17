package assessment

type Assessed struct {
	AssdImprValue float32
	AssdLandValue float32
	AssdTtlValue  float32
}

func NewAssessed(assdImprValue float32, assdLandValue float32, assdTtlValue float32) Assessed {
	return Assessed{
		AssdImprValue: assdImprValue,
		AssdLandValue: assdLandValue,
		AssdTtlValue:  assdTtlValue,
	}
}
