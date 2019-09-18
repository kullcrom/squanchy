package sale

//SaleAmountData ..
type SaleAmountData struct {
	SaleAmt            float32 `json:"saleAmt"`
	SaleDisclosureType int     `json:"saleDisclosureType"`
}

//NewSaleAmountData ...
func NewSaleAmountData(saleAmt float32, saleDisclosureType int) SaleAmountData {
	return SaleAmountData{
		SaleAmt:            saleAmt,
		SaleDisclosureType: saleDisclosureType,
	}
}
