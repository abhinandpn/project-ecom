package repository

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.Cartrepository {

	return &cartDatabase{DB: db}
}

func (c *cartDatabase) FindCartByUID(uid uint) (domain.Cart, error) {

	var body domain.Cart

	query := `select * from carts where user_id = $1`
	err := c.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *cartDatabase) FindCartInfoByCID(cid uint) (domain.CartInfo, error) {

	var body domain.CartInfo

	query := `select * from cart_infos where cart_id = $1`
	err := c.DB.Raw(query, cid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (c *cartDatabase) FindProductFromCartByCId(pid uint) (domain.Cart, error) {

	var body domain.Cart

	query := `select * from carts where product_id = $1`
	err := c.DB.Raw(query, pid).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *cartDatabase) FindProductFromCartInfoByCId(pid uint) (domain.CartInfo, error) {

	var body domain.CartInfo

	query := `select * from cart_infos where product_id = $1`
	err := c.DB.Raw(query, pid).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *cartDatabase) CreateCartByUID(uid uint) (domain.Cart, error) {

	var body domain.Cart
	query := `insert into carts (user_id)values ($1);`
	err := c.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *cartDatabase) CreateCartInfoByCid(cid uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	query := `insert into cart_infos (cart_id)values ($1);`
	err := c.DB.Raw(query, cid).Error
	if err != nil {
		return body, err
	}

	return body, nil
}

func (c *cartDatabase) AddProductToCart(uid, pid, pfid uint) error {

	query := `insert into carts (user_id,product_id,product_info,quantity)values ($1,$2,$3,$4);`
	err := c.DB.Raw(query, uid, pid, pfid).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cartDatabase) AddProductToCartInfo(cid uint, pfr domain.Product) error {

	query := `insert into cart_infos (cart_id,sub_total)values ($1,$2);`
	err := c.DB.Raw(query, cid, pfr.Price).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cartDatabase) RemoveProductfromCart(uid, pfid uint) error {

	query := `delete * from carts where user_id = $1 and product_info = $2;`
	err := c.DB.Raw(query, uid, pfid).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cartDatabase) RemoveProductfromCartInfo(cid uint) error {

	query := `delete * from cart_infos where cart_id = $1 ;`
	err := c.DB.Raw(query, cid).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *cartDatabase) ListAllProductFromCart(pagination req.PageNation, uid uint) ([]res.DisplayCart, error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit
	var body []res.DisplayCart
	query := `SELECT
					p.product_name,
					p.image,
					pi.colour,
					pi.size,
					ci.total,
					c.quantity,
					ci.total + c.quantity AS total_amount
 		FROM carts c
 		INNER JOIN product_infos pi ON pi.product_id = c.product_id
 		INNER JOIN products p ON p.id = pi.product_id
 		INNER JOIN cart_infos ci ON ci.cart_id = c.id
 		WHERE c.user_id = $1
		 LIMIT
		 	$2 OFFSET $3;`
	err := c.DB.Raw(query, uid, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
