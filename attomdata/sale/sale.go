package sale

type Sale struct {
	SequenceSaleHistory int
	SaleAmountData      SaleAmountData
}

func NewSale(sequenceSaleHistory int, saleAmountData SaleAmountData) Sale {
	return Sale{
		SequenceSaleHistory: sequenceSaleHistory,
		SaleAmountData:      saleAmountData,
	}
}
