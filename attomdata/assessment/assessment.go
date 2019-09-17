package assessment

import (
	"github.com/kullcrom/squanchy/attomdata/mortgage"
	"github.com/kullcrom/squanchy/attomdata/owner"
	"github.com/kullcrom/squanchy/attomdata/tax"
)

type Assessment struct {
	Appraised          Appraised
	Assessed           Assessed
	Market             Market
	Tax                tax.Tax
	DelinquentYear     int
	ImprovementPercent int
	Owner              owner.Owner
	Mortgage           mortgage.Mortgage
}
