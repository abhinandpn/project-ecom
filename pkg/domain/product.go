package domain

import "time"

type Catogery struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	CategoryName string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}
type Product struct {
	Id              uint   `gorm:"primaryKey;unique;not null"`
	ProductName     string `gorm:"unique;not null"`
	Description     string
	Brand           string
	Category_id     uint
	Category        Catogery `gorm:"foreignKey:Category_id"`
	Price           uint
	Size            uint
	Colour          string
	Img             string
	SKU             string
	QuantityInStock uint
	Created_at      time.Time
	Updated_at      time.Time
}
type ProductItem struct {
	Id         uint
	Product_Id uint
	Product    Product
}
type Images struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	ProductItemId uint
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemId"`
	FileName      string
}
type ProductBrand struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Brand string `gorm:"not null,index,unique" json:"brand" validate:"required"`
}
