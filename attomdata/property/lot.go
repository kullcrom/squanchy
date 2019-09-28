package property

//Lot ...
type Lot struct {
	Depth      float32 `json:"depth"`
	Frontage   float32 `json:"frontage"`
	LotSize1   float32 `json:"lotSize1"`
	LotSize2   float32 `json:"lotSize2"`
	ZoningType string  `json:"zoningType"`
}

//NewLot ...
func NewLot(
	depth float32,
	frontage float32,
	lotSize1 float32,
	lotSize2 float32,
	zoningType string) *Lot {
	lot := Lot{
		Depth:      depth,
		Frontage:   frontage,
		LotSize1:   lotSize1,
		LotSize2:   lotSize2,
		ZoningType: zoningType,
	}
	return &lot
}
