package SalseTracker

// type Response struct {
// 	OrderID         string `json:"OrderID"`
// 	ProductID       string `json:"ProductID"`
// 	CustomerID      string `json:"CustomerID"`
// 	ProductName     string `json:"ProductName"`
// 	Category        string `json:"Category"`
// 	Region          string `json:"Region"`
// 	DateofSale      string `json:"DateofSale"`
// 	QuantitySold    string `json:"QuantitySold"`
// 	UnitPrice       string `json:"UnitPrice"`
// 	Discount        string `json:"Discount"`
// 	ShippingCost    string `json:"ShippingCost"`
// 	PaymentMethod   string `json:"PaymentMethod"`
// 	CustomerName    string `json:"CustomerName"`
// 	CustomerEmail   string `json:"CustomerEmail"`
// 	CustomerAddress string `json:"CustomerAddress"`
// }

type ProductMaster struct {
	ProductID       string `json:"ProductID"`
	ProductName     string `json:"ProductName"`
	Category        string `json:"Category"`
}

type Sales struct {
	OrderID       string `json:"OrderID"`
	DateofSale    string `json:"DateofSale"`
	QuantitySold  float64 `json:"QuantitySold"`
	UnitPrice     float64 `json:"UnitPrice"`
	Discount      float64 `json:"Discount"`
	ShippingCost  float64 `json:"ShippingCost"`
	PaymentMethod string `json:"PaymentMethod"`
}

type CoustomerDetalis struct {
	CustomerID      string `json:"CustomerID"`
	CustomerName    string `json:"CustomerName"`
	CustomerEmail   string `json:"CustomerEmail"`
	CustomerAddress string `json:"CustomerAddress"`
	Region          string `json:"Region"`
}