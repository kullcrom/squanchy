package owner

//Owner ...
type Owner struct {
	Owner1                Owner1 `json:"owner1"`
	Owner2                Owner2 `json:"owner2"`
	AbsenteeOwnerStatus   string `json:"absenteeOwnerStatus"`
	MailingAddressOneLine string `json:"mailingAddressOneLine"`
}

//NewOwner ...
func NewOwner(owner1 Owner1, owner2 Owner2, absenteeOwnerStatus string, mailingAddressOneLine string) Owner {
	return Owner{
		Owner1:                owner1,
		Owner2:                owner2,
		AbsenteeOwnerStatus:   absenteeOwnerStatus,
		MailingAddressOneLine: mailingAddressOneLine,
	}
}
