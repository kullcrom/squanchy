package property

import (
	"github.com/kullcrom/squanchy/attomdata/assessment"
	"github.com/kullcrom/squanchy/attomdata/building"
	"github.com/kullcrom/squanchy/attomdata/sale"
	"github.com/kullcrom/squanchy/attomdata/vintage"
)

//Property ...
type Property struct {
	Identifier Identifier            `json:"identifier"`
	Lot        Lot                   `json:"lot"`
	Area       Area                  `json:"area"`
	Address    Address               `json:"address"`
	Location   Location              `json:"location"`
	Summary    Summary               `json:"summary"`
	Utilities  Utilities             `json:"utilities"`
	Sale       sale.Sale             `json:"sale"`
	Building   building.Building     `json:"building"`
	Assessment assessment.Assessment `json:"assessment"`
	Vintage    vintage.Vintage       `json:"vintage"`
}

//NewProperty ...
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
