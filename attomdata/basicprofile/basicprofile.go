package basicprofile

import (
	"github.com/kullcrom/squanchy/attomdata/property"
)

//BasicProfile ...
type BasicProfile struct {
	Status       Status              `json:"status"`
	EchoedFields EchoedFields        `json:"echoed_fields"`
	Property     []property.Property `json:"property"`
}

//NewBasicProfile ...
func NewBasicProfile(status Status, echoedFields EchoedFields, property []property.Property) BasicProfile {
	return BasicProfile{
		Status:       status,
		EchoedFields: echoedFields,
		Property:     property,
	}
}
