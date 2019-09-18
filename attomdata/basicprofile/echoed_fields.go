package basicprofile

type EchoedFields struct {
	JobID       string `json:"jobID"`
	LoanNumber  string `json:"loanNumber"`
	PreparedBy  string `json:"preparedBy"`
	ResellerID  string `json:"resellerID"`
	PreparedFor string `json:"preparedFor"`
}

func NewEchoedFields(
	jobID string,
	loanNumber string,
	preparedBy string,
	resellerID string,
	preparedFor string) EchoedFields {
	return EchoedFields{
		JobID:       jobID,
		LoanNumber:  loanNumber,
		PreparedBy:  preparedBy,
		ResellerID:  resellerID,
		PreparedFor: preparedFor,
	}
}
