package mortgage

//FirstConcurrent ...
type FirstConcurrent struct {
	Amount         float32 `json:"amount"`
	LenderLastName string  `json:"lenderLastName"`
}

//NewFistConcurrent ...
func NewFirstConcurrent(amount float32, lenderLastName string) FirstConcurrent {
	return FirstConcurrent{
		Amount:         amount,
		LenderLastName: lenderLastName,
	}
}
