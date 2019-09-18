package tax

//Tax ...
type Tax struct {
	TaxAmt         float32   `json:"taxAmt"`
	TaxPerSizeUnit float32   `json:"taxPerSizeUnit"`
	TaxYear        float32   `json:"taxYear"`
	Exemption      Exemption `json:"exemption"`
}

//NewTax ...
func NewTax(taxAmt float32, taxPerSizeUnit float32, taxYear float32, exemption Exemption) Tax {
	return Tax{
		TaxAmt:         taxAmt,
		TaxPerSizeUnit: taxPerSizeUnit,
		TaxYear:        taxYear,
		Exemption:      exemption,
	}
}
