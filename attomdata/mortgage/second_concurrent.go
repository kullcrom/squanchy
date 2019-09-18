package mortgage

//SecondConcurrent ...
type SecondConcurrent struct {
	Amount float32 `json:"amount"`
}

//NewSecondConcurrent ...
func NewSecondConcurrent(amount float32) SecondConcurrent {
	return SecondConcurrent{
		Amount: amount,
	}
}
