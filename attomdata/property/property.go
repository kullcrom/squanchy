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

func NewProperty(
	identifier Identifier,
	lot Lot,
	area Area,
	address Address,
	location Location,
	summary Summary,
	utilities Utilities,
	sale sale.Sale,
	building building.Building,
	assessment assessment.Assessment,
	vintage vintage.Vintage) Property {
	return Property{
		Identifier: identifier,
		Lot:        lot,
		Area:       area,
		Address:    address,
		Location:   location,
		Summary:    summary,
		Utilities:  utilities,
		Sale:       sale,
		Building:   building,
		Assessment: assessment,
		Vintage:    vintage,
	}
}
