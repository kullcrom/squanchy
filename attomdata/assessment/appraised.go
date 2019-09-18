package assessment

//Appraised ...
type Appraised struct {
	ApprImprValue float32 `json:"apprImprValue"`
	ApprLandValue float32 `json:"apprLandValue"`
	ApprTTLValue  float32 `json:"apprTtlValue"`
}

//NewAppraised ...
func NewAppraised(apprImprValue float32, apprLandValue float32, apprTTLValue float32) Appraised {
	return Appraised{
		ApprImprValue: apprImprValue,
		ApprLandValue: apprLandValue,
		ApprTTLValue:  apprTTLValue,
	}
}
