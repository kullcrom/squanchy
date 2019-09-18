package mortgage

//Mortgage ...
type Mortgage struct {
	FirstConcurrent  FirstConcurrent  `json:"FirstConcurrent"`
	SecondConcurrent SecondConcurrent `json:"SecondConcurrent"`
	ThirdConcurrent  ThirdConcurrent  `json:"ThirdConcurrent"`
}

//NewMortgage ...
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
