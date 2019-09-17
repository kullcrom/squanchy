package property

type Address struct {
	Country     string
	CountrySubd string
	Line1       string
	Line2       string
	Locality    string
	MatchCode   string
	OneLine     string
	Postal1     string
	Postal2     string
	Postal3     string
}

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
