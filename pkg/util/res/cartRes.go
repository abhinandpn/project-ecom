package res

type CartDisplay struct {
	Id           uint
	ProductName  string
	Size         uint
	Colour       string
	BrandName    string
	CategoryName string
	Price        float64
	Quantity     uint
}

type CartInfo struct {
	Subtotal      float64
	DIscountPrice float64
	CouponCode    string
	Totalprice    float64
}

type ProductInfoCart struct {
	ProductInfoID int `json:"product_info_id"`
}
