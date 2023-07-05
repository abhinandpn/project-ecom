package res

type CartDisplay struct {
	Id           uint
	ProductName  string
	Size         uint
	Colour       string
	BrandName    string
	CategoryName string
	Price        float64
	// Quantity     uint
}

type CartInfo struct {
	Subtotal      float64
	DIscountPrice float64
	Totalprice    float64
}
