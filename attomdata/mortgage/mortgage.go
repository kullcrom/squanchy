package mortgage

type Mortgage struct {
	FirstConcurrent  FirstConcurrent
	SecondConcurrent SecondConcurrent
	ThirdConcurrent  ThirdConcurrent
}

func NewMortgage(
	firstConcurrent FirstConcurrent,
	secondConcurrent SecondConcurrent,
	thirdConcurrent ThirdConcurrent) Mortgage {
	return Mortgage{
		FirstConcurrent:  firstConcurrent,
		SecondConcurrent: secondConcurrent,
		ThirdConcurrent:  thirdConcurrent,
	}
}
