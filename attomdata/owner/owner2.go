package owner

//Owner2 ...
type Owner2 struct {
	LastName       string `json:"lastName"`
	FirstNameAndMI string `json:"firstNameAndMi"`
}

//NewOwner2 ...
func NewOwner2(lastName string, firstNameAndMI string) Owner2 {
	return Owner2{
		LastName:       lastName,
		FirstNameAndMI: firstNameAndMI,
	}
}
