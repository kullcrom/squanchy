package basicprofile

type Status struct {
	Version          string `json:"version"`
	Code             int    `json:"code"`
	Msg              string `json:"msg"`
	Total            int    `json:"total"`
	Page             int    `json:"page"`
	PageSize         int    `json:"pagesize"`
	ResponseDateTime string `json:"responseDateTime"`
	TransactionID    string `json:"transactionID"`
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
