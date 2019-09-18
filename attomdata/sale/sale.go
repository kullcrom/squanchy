package sale

//Sale ...
type Sale struct {
	SequenceSaleHistory int            `json:"sequenceSaleHistory"`
	SaleAmountData      SaleAmountData `json:"saleAmountData"`
}

//NewSale ...
func NewSale(sequenceSaleHistory int, saleAmountData SaleAmountData) Sale {
	return Sale{
		SequenceSaleHistory: sequenceSaleHistory,
		SaleAmountData:      saleAmountData,
	}
}
