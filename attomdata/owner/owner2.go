package owner

type Owner2 struct {
	LastName       string
	FirstNameAndMI string
}

func NewOwner2(lastName string, firstNameAndMI string) Owner2 {
	return Owner2{
		LastName:       lastName,
		FirstNameAndMI: firstNameAndMI,
	}
}
