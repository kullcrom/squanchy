package building

type Summary struct {
	Levels     int
	StoryDesc  string
	UnitsCount int
	ViewCode   string
}

func NewSummary(levels int, storyDesc string, unitsCount int, viewCode string) Summary {
	return Summary{
		Levels:     levels,
		StoryDesc:  storyDesc,
		UnitsCount: unitsCount,
		ViewCode:   viewCode,
	}
}
