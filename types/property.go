package types

//Property ...
type Property struct {
	ID            int
	UserID        int
	street        string
	city          string
	state         string
	PurchasePrice float32
	ExitStrategy  int
}
