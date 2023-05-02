package domain

type Favourite struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	UserId      uint
	Users       Users `gorm:"foreignKey:UserId"`
	ItemId      uint
	ProductItem ProductItem `gorm:"foreignKey:ItemId"`
}
