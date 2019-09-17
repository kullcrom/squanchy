package basicprofile

import (
	"github.com/kullcrom/squanchy/attomdata/property"
)

type BasicProfile struct {
	Status       Status
	EchoedFields EchoedFields
	Property     property.Property
}

func NewBasicProfile(status Status, echoedFields EchoedFields, property property.Property) BasicProfile {
	return BasicProfile{
		Status:       status,
		EchoedFields: echoedFields,
		Property:     property,
	}
}
