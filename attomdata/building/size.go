package building

type Size struct {
	BldgSize          float32
	GrossSize         float32
	GrossSizeAdjusted float32
	GroundFloorSize   float32
	LivingSize        float32
	SizeInd           string
	UniversalSize     float32
	AtticSize         float32
}

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
