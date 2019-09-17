package sale

type SaleAmountData struct {
	SaleAmt            float32
	SaleDisclosureType int
}

func NewSaleAmountData(saleAmt float32, saleDisclosureType int) SaleAmountData {
	return SaleAmountData{
		SaleAmt:            saleAmt,
		SaleDisclosureType: saleDisclosureType,
	}
}
