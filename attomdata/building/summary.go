package building

//Summary ...
type Summary struct {
	Levels     int    `json:"levels"`
	StoryDesc  string `json:"storyDesc"`
	UnitsCount int    `json:"unitsCount"`
	ViewCode   string `json:"viewCode"`
}

//NewSummary ...
func NewSummary(levels int, storyDesc string, unitsCount int, viewCode string) Summary {
	return Summary{
		Levels:     levels,
		StoryDesc:  storyDesc,
		UnitsCount: unitsCount,
		ViewCode:   viewCode,
	}
}
