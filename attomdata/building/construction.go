package building

type Construction struct {
	condition string
}

func GetCondition() string {
	return this.condition
}

func SetCondition(condition string) {
	this.condition = condition
}

func NewConstruction(condition string) Construction {
	this.SetCondition(condition)
	return this
}
