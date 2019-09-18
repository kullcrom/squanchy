package building

//Size ...
type Size struct {
	BldgSize          float32 `json:"bldgSize"`
	GrossSize         float32 `json:"grossSize"`
	GrossSizeAdjusted float32 `json:"grossSizeAdjusted"`
	GroundFloorSize   float32 `json:"groundFloorSize"`
	LivingSize        float32 `json:"livingSize"`
	SizeInd           string  `json:"sizeInd"`
	UniversalSize     float32 `json:"universalSize"`
	AtticSize         float32 `json:"atticSize"`
}

//NewSize ...
func NewSize(
	bldgSize float32,
	grossSize float32,
	grossSizeAdjusted float32,
	groundFloorSize float32,
	livingSize float32,
	sizeInd string,
	universalSize float32,
	atticSize float32) Size {
	return Size{
		BldgSize:          bldgSize,
		GrossSize:         grossSize,
		GrossSizeAdjusted: grossSizeAdjusted,
		GroundFloorSize:   groundFloorSize,
		LivingSize:        livingSize,
		SizeInd:           sizeInd,
		UniversalSize:     universalSize,
		AtticSize:         atticSize,
	}
}
