package property

import (
	"github.com/kullcrom/squanchy/attomdata/assessment"
	"github.com/kullcrom/squanchy/attomdata/building"
	"github.com/kullcrom/squanchy/attomdata/sale"
	"github.com/kullcrom/squanchy/attomdata/vintage"
)

type Property struct {
	Identifier Identifier
	Lot        Lot
	Area       Area
	Address    Address
	Location   Location
	Summary    Summary
	Utilities  Utilities
	Sale       sale.Sale
	Building   building.Building
	Assessment assessment.Assessment
	Vintage    vintage.Vintage
}
