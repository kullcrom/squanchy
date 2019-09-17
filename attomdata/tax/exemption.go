package tax

type Exemption struct {
	ExemptionAccount1 float32
	ExemptionAccount2 float32
	ExemptionAccount3 float32
	ExemptionAccount4 float32
	ExemptionAccount5 float32
}

func NewExemption(
	exemptionAccount1 float32,
	exemptionAccount2 float32,
	exemptionAccount3 float32,
	exemptionAccount4 float32,
	exemptionAccount5 float32,
) Exemption {
	return Exemption{
		ExemptionAccount1: exemptionAccount1,
		ExemptionAccount2: exemptionAccount2,
		ExemptionAccount3: exemptionAccount3,
		ExemptionAccount4: exemptionAccount4,
		ExemptionAccount5: exemptionAccount5,
	}
}
