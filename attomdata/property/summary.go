package property

//Summary ...
type Summary struct {
	AbsenteeInd   string `json:"absenteeInd"`
	PropClass     string `json:"propClass"`
	PropSubType   string `json:"propSubType"`
	PropType      string `json:"propType"`
	YearBuilt     int    `json:"yearBuilt"`
	PropLandUse   string `json:"propLandUse"`
	PropIndicator int    `json:"propIndicator"`
	Legal1        string `json:"legal1"`
}

//NewSummary ...
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
