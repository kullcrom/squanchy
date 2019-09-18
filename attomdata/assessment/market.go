package assessment

//Market ...
type Market struct {
	MktImprValue float32 `json:"mktImprValue"`
	MktLandValue float32 `json:"mktLandValue"`
	MktTTLValue  float32 `json:"mktTtlValue"`
}

//NewMarket ...
func NewMarket(mktImprValue float32, mktLandValue float32, mktTTLValue float32) Market {
	return Market{
		MktImprValue: mktImprValue,
		MktLandValue: mktLandValue,
		MktTTLValue:  mktTTLValue,
	}
}
