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
