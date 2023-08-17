package res

import "github.com/abhinandpn/project-ecom/pkg/domain"

type OrderStatus struct {
	Name        string
	Adress      domain.Address
	OrderType   string
	OrderStatus string
	Products    ResProduct
	CartRes     []CartDisplay
}

type OrderStatusRes struct {
	Address       []domain.Address
	Product       []ResProduct
	TotalPrice    float64
	PaymentMethod string
	OrderStatus   string
}
type ResOrder struct {
	Id            uint
	AddressId     uint
	Total_price   float64
	PaymentId     uint
	ProductInfoId uint
	Quantity      uint
	OrderStatus   string
}
type UpdateOrderDetail struct {
	Address       []domain.Address
	PaymentMethod string
	PaymentStatus string
	TotalAmmount  float64
	Product       []ResProduct
	OrderStatus   string
}
