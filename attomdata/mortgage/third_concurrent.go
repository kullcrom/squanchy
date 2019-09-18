package mortgage

//ThirdConcurrent ...
type ThirdConcurrent struct {
	Amount float32 `json:"amount"`
}

//NewThirdConcurrent ...
func NewThirdConcurrent(amount float32) ThirdConcurrent {
	return ThirdConcurrent{
		Amount: amount,
	}
}
