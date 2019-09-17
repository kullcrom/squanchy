package owner

type Owner1 struct {
	LastName       string
	FirstNameAndMI string
}

func NewOwner1(lastName string, firstNameAndMI string) Owner1 {
	return Owner1{
		LastName:       lastName,
		FirstNameAndMI: firstNameAndMI,
	}
}
