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
func (c *cartDatabase) FindCartByUId(id uint) (domain.UserCart, error) {

	var body domain.UserCart
	query := `select * from user_carts where user_id = $1`
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
	query := `insert into cart_infos (cart_id,product_info_id,quantity)values ($1,$2,$3);`

	err := c.DB.Raw(query, id, pfid, 1).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartDatabase) RemoveCart(id, pfid uint) error {

	query := `DELETE FROM cart_infos
				WHERE cart_id = $1 AND product_info_id = $2;	`
	err := c.DB.Exec(query, id, pfid).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cartDatabase) ViewCart(id uint) ([]res.CartDisplay, error) {

	
	var body []res.CartDisplay
	query := `SELECT p.product_name, pi.size, pi.colour, b.brand_name, c.category_name, pi.price
					FROM user_carts uc
					JOIN cart_infos ci ON uc.id = ci.cart_id
					JOIN products p ON ci.product_info_id = p.id
					JOIN product_infos pi ON p.id = pi.product_id
					JOIN brands b ON p.brand_id = b.id
					JOIN categories c ON p.category_id = c.id
					WHERE uc.user_id = $1;`
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

func (c *cartDatabase) CartInfo(id uint) (res.CartInfo, error) {

	var body res.CartInfo

	query := `SELECT
					SUM(pi.price * ci.quantity) AS subtotal,
					0 AS discountprice,
					SUM(pi.price * ci.quantity) - 0 AS totalprice
				  FROM
					user_carts uc
				  JOIN
					cart_infos ci ON uc.id = ci.cart_id
				  JOIN
					product_infos pi ON ci.product_info_id = pi.id
				  WHERE
					uc.user_id = $1;`
	err := c.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
