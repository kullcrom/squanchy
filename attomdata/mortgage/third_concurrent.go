package mortgage

type ThirdConcurrent struct {
	Amount float32
}

func NewThirdConcurrent(amount float32) ThirdConcurrent {
	return ThirdConcurrent{
		Amount: amount,
	}
}
