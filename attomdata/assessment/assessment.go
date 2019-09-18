package assessment

import (
	"github.com/kullcrom/squanchy/attomdata/mortgage"
	"github.com/kullcrom/squanchy/attomdata/owner"
	"github.com/kullcrom/squanchy/attomdata/tax"
)

//Assessment ...
type Assessment struct {
	Appraised          Appraised         `json:"appraised"`
	Assessed           Assessed          `json:"assessed"`
	Market             Market            `json:"market"`
	Tax                tax.Tax           `json:"tax"`
	DelinquentYear     int               `json:"delinquentyear"`
	ImprovementPercent int               `json:"improvementPercent"`
	Owner              owner.Owner       `json:"owner"`
	Mortgage           mortgage.Mortgage `json:"mortgage"`
}
