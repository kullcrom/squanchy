package basicprofile

type Status struct {
	Version          string
	Code             int
	Msg              string
	Total            int
	Page             int
	PageSize         int
	ResponseDateTime string
	TransactionID    string
}

func NewStatus(
	version string,
	code int,
	msg string,
	total int,
	page int,
	pageSize int,
	responseDateTime string,
	transactionID string) Status {
	return Status{
		Version:          version,
		Code:             code,
		Msg:              msg,
		Total:            total,
		Page:             page,
		PageSize:         pageSize,
		ResponseDateTime: responseDateTime,
		TransactionID:    transactionID,
	}
}
