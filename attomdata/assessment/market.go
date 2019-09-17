package assessment

type Market struct {
	MktImprValue float32
	MktLandValue float32
	MktTtlValue  float32
}

func NewMarket(mktImprValue float32, mktLandValue float32, mktTtlValue float32) Market {
	return Market{
		MktImprValue: mktImprValue,
		MktLandValue: mktLandValue,
		MktTtlValue:  mktTtlValue,
	}
}
