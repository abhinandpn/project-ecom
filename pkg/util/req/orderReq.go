package req

type UserOrderReq struct {
	AddressId       uint `json:"address_id"`
	PaymentMethodId uint `json:"payment_method_id"`
	ProductInfoId   uint `json:"product_info_id"`
	Quantity        uint `json:"quantity"`
}
