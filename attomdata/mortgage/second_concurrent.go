package mortgage

type SecondConcurrent struct {
	Amount float32
}

func NewSecondConcurrent(amount float32) SecondConcurrent {
	return SecondConcurrent{
		Amount: amount,
	}
}
