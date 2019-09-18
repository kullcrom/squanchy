package tax

//Exemption ...
type Exemption struct {
	ExemptionAmount1 float32 `json:"ExemptionAmount1"`
	ExemptionAmount2 float32 `json:"ExemptionAmount2"`
	ExemptionAmount3 float32 `json:"ExemptionAmount3"`
	ExemptionAmount4 float32 `json:"ExemptionAmount4"`
	ExemptionAmount5 float32 `json:"ExemptionAmount5"`
}

//NewExemption ...
func NewExemption(
	exemptionAmount1 float32,
	exemptionAmount2 float32,
	exemptionAmount3 float32,
	exemptionAmount4 float32,
	exemptionAmount5 float32,
) Exemption {
	return Exemption{
		ExemptionAmount1: exemptionAmount1,
		ExemptionAmount2: exemptionAmount2,
		ExemptionAmount3: exemptionAmount3,
		ExemptionAmount4: exemptionAmount4,
		ExemptionAmount5: exemptionAmount5,
	}
}
