package owner

type Owner struct {
	Owner1                Owner1
	Owner2                Owner2
	AbsenteeOwnerStatus   string
	MailingAddressOneLine string
}

func NewOwner(owner1 Owner1, owner2 Owner2, absenteeOwnerStatus string, mailingAddressOneLine string) Owner {
	return Owner{
		Owner1:                owner1,
		Owner2:                owner2,
		AbsenteeOwnerStatus:   absenteeOwnerStatus,
		MailingAddressOneLine: mailingAddressOneLine,
	}
}
