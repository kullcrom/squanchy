package basicprofile

type EchoedFields struct {
	JobID       string
	LoanNumber  string
	PreparedBy  string
	ResellerID  string
	PreparedFor string
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
