package property

type Summary struct {
	AbsenteeInd   string
	PropClass     string
	PropSubType   string
	PropType      string
	YearBuilt     int
	PropLandUse   string
	PropIndicator int
	Legal1        string
}

func NewSummary(
	absenteeInd string,
	propClass string,
	propSubType string,
	propType string,
	yearBuilt int,
	propLandUse string,
	propIndicator int,
	legal1 string) Summary {
	return Summary{
		AbsenteeInd:   absenteeInd,
		PropClass:     propClass,
		PropSubType:   propSubType,
		PropType:      propType,
		YearBuilt:     yearBuilt,
		PropLandUse:   propLandUse,
		PropIndicator: propIndicator,
		Legal1:        legal1,
	}
}
