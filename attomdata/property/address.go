package property

//Address ...
type Address struct {
	Country     string `json:"country"`
	CountrySubd string `json:"countrySubd"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2"`
	Locality    string `json:"locality"`
	MatchCode   string `json:"matchCode"`
	OneLine     string `json:"oneLine"`
	Postal1     string `json:"postal1"`
	Postal2     string `json:"postal2"`
	Postal3     string `json:"postal3"`
}

//NewAddress ...
func NewAddress(
	country string,
	countrySubd string,
	line1 string,
	line2 string,
	locality string,
	matchCode string,
	oneLine string,
	postal1 string,
	postal2 string,
	postal3 string) Address {
	return Address{
		Country:     country,
		CountrySubd: countrySubd,
		Line1:       line1,
		Line2:       line2,
		Locality:    locality,
		MatchCode:   matchCode,
		OneLine:     oneLine,
		Postal1:     postal1,
		Postal2:     postal2,
		Postal3:     postal3,
	}
}
