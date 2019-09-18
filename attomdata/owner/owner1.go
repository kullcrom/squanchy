package owner

//Owner1 ...
type Owner1 struct {
	LastName       string `json:"lastName"`
	FirstNameAndMI string `json:"firstNameAndMi"`
}

//NewOwner1 ...
func NewOwner1(lastName string, firstNameAndMI string) Owner1 {
	return Owner1{
		LastName:       lastName,
		FirstNameAndMI: firstNameAndMI,
	}
}
