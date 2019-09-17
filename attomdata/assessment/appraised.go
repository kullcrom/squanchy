package assessment

type Appraised struct {
	ApprImprValue float32
	ApprLandValue float32
	ApprTtlValue  float32
}

func NewAppraised(apprImprValue float32, apprLandValue float32, apprTtlValue float32) Appraised {
	return Appraised{
		ApprImprValue: apprImprValue,
		ApprLandValue: apprLandValue,
		ApprTtlValue:  apprTtlValue,
	}
}
