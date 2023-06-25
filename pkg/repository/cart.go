package repository

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.Cartrepository {

	return &cartDatabase{DB: db}
}
func (c *cartDatabase) FindCartBy(id uint) (domain.UserCart, error) {

	var body domain.UserCart
	query := `select * from user_carts where id = $1`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *cartDatabase) FindCartInfoById(id uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	query := `select * from cart_infos where id =$1;`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *cartDatabase) CreateUserCart(id uint) (domain.UserCart, error) {

	var body domain.UserCart
	query := `insert into user_carts (user_id)values ($1) returning id;`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *cartDatabase) CreateCartinfo(id uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	query := `insert into cart_infos (cart_id)values($1) returning id;`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *cartDatabase) AddToCart(id, pfid, qty uint) error {

	var body domain.CartInfo
	query := `UPDATE cart_infos
					SET product_info_id = $1,
					  quantity = $2
					WHERE cart_id = $3;`

	err := c.DB.Raw(query, pfid, qty, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartDatabase) RemoveCart(id, pfid uint) error {

	query := `delete from cart_infos where cart_id = $1;`
	err := c.DB.Exec(query, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartDatabase) ViewCart(id uint) (res.CartDisplay, error) {

	var body res.CartDisplay
	query := `SELECT
				    p.product_name,
				    pi.size,
				    pi.colour,
				    b.brand_name,
				    c.category_name,
				    pi.price
				FROM
				    user_carts uc
				    JOIN cart_infos ci ON uc.id = ci.cart_id
				    JOIN product_infos pi ON ci.product_info_id = pi.id
				    JOIN products p ON pi.product_id = p.id
				    JOIN brands b ON p.brand_id = b.id
				    JOIN categories c ON p.category_id = c.id
				WHERE
				    uc.id = $1;`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, err

}

func (c *cartDatabase) FindProductIntoCart(id, pfid uint) (bool, error) {

	var body bool
	query := `select  exists(select * from cart_infos where cart_id = $1 and product_info_id = $2);`
	err := c.DB.Raw(query, id, pfid).Scan(&body).Error
	if err != nil {

		return body, err
	}
	if !body {
		return body, nil
	}
	return body, nil

}

// func (c *cartDatabase) CartInfo(id uint) (res.CartInfo, error) {

// 	var body res.CartInfo

// }
