package res

import (
	"time"

	"github.com/abhinandpn/project-ecom/pkg/domain"
)

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
type UpdatedOrderInfo struct {
	OrderId       uint
	OrderTime     time.Time
	Address       domain.Address
	CouponCode    string
	TotalPrice    float64
	OrderStatus   string
	PayentMethod  string
	PaymentDetail string
}

type OrderDetailByUid struct {
	OrderInfoId uint
	OrderId     uint
	OrderTime   time.Time
	// Addressid      uint
	House          string
	PhoneNumber    string
	Street         string
	City           string
	District       string
	Pincode        uint
	Landmark       string
	AddressName    string
	CouponCode     uint
	CouponName     string
	CouponDiscount uint
	TotalPrice     float64
	OrderStatus    string
	// PaymenentTotalPrice float64
	PaymentMethod string
	PaymentStatus string
}
