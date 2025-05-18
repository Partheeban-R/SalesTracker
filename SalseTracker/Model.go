package SalseTracker

type ProductMaster struct {
	ProductID   string `json:"ProductID"`
	ProductName string `json:"ProductName"`
	Category    string `json:"Category"`
}

type Sales struct {
	OrderID       string  `json:"OrderID"`
	DateofSale    string  `json:"DateofSale"`
	QuantitySold  float64 `json:"QuantitySold"`
	UnitPrice     float64 `json:"UnitPrice"`
	Discount      float64 `json:"Discount"`
	ShippingCost  float64 `json:"ShippingCost"`
	PaymentMethod string  `json:"PaymentMethod"`
}

type CoustomerDetalis struct {
	CustomerID      string `json:"CustomerID"`
	CustomerName    string `json:"CustomerName"`
	CustomerEmail   string `json:"CustomerEmail"`
	CustomerAddress string `json:"CustomerAddress"`
	Region          string `json:"Region"`
}

type Revenue struct {
	Product  string  `json:"Product,omitempty"`
	Revenue  float64 `json:"Revenue,omitempty"`
	Category string  `json:"Category,omitempty"`
	Region   string  `json:"Region,omitempty"`
}

type RevenueResp struct {
	TotalRevenue    float64   `json:"TotalRevenue,omitempty"`
	ProductRevenue  []Revenue `json:"ProductRevenue,omitempty"`
	CategoryRevenue []Revenue `json:"CategoryRevenue,omitempty"`
	RegionRevenue   []Revenue `json:"RegionRevenue,omitempty"`
	Status          string    `json:"status"`
	ErrMsg          string    `json:"errMsg"`
}
