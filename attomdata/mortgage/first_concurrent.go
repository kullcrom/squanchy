package mortgage

type FirstConcurrent struct {
	Amount         float32
	LenderLastName string
}

func NewFirstConcurrent(amount float32, lenderLastName string) FirstConcurrent {
	return FirstConcurrent{
		Amount:         amount,
		LenderLastName: lenderLastName,
	}
}
