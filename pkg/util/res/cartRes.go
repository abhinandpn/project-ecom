package res

type CartDisplay struct {
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
	TotalPrice    float64
}
