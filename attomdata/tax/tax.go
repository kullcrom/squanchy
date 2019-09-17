package tax

type Tax struct {
	TaxAmt         float32
	TaxPerSizeUnit float32
	TaxYear        float32
	Exemption      Exemption
}

func NewTax(taxAmt float32, taxPerSizeUnit float32, taxYear float32, exemption Exemption) Tax {
	return Tax{
		TaxAmt:         taxAmt,
		TaxPerSizeUnit: taxPerSizeUnit,
		TaxYear:        taxYear,
		Exemption:      exemption,
	}
}
