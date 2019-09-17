package property

type Lot struct {
	Depth      float32
	Frontage   float32
	LotSize1   float32
	LotSize2   float32
	ZoningType string
}

func NewLot(depth float32, frontage float32, lotSize1 float32, lotSize2 float32, zoningType string) Lot {
	return Lot{
		Depth:      depth,
		Frontage:   frontage,
		LotSize1:   lotSize1,
		LotSize2:   lotSize2,
		ZoningType: zoningType,
	}
}
