package building

type Construction struct {
	Condition string
}

func NewConstruction(condition string) Construction {
	return Construction{Condition: condition}
}
