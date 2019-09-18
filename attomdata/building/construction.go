package building

//Construction ...
type Construction struct {
	Condition string `json:"condition"`
}

//NewConstruction ...
func NewConstruction(condition string) Construction {
	return Construction{Condition: condition}
}
