package repository

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.Cartrepository {

	return &cartDatabase{DB: db}
}

// function for create a empty cart for user

func (crt *cartDatabase) CreateCart(ctx context.Context, uid uint) (domain.Cart, error) {

	var body domain.Cart
	query := `INSERT INTO carts (user_id,total_price) VALUES($1, $2) RETURNING cart_id`

	err := crt.DB.Raw(query, uid, 0).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (crt *cartDatabase) FindCartByUserId(ctx context.Context, uid uint) (domain.Cart, error) {

	var body domain.Cart

	query := `select * from carts where user_id = ?;`
	err := crt.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (crt *cartDatabase) FindCartIdByUserId(ctx context.Context, id uint) (uint, error) {

	var Cartid uint
	query := ` select id from carts where id = ?;`
	err := crt.DB.Raw(query, id).Error
	if err != nil {
		return Cartid, err
	}
	return Cartid, nil
}

func (crt *cartDatabase) UpdateCartHelp(ctx context.Context, uid uint, price float64) error {

	var cart domain.Cart
	query := `update carts set total_price =$1 where user_id = $2;`
	err := crt.DB.Raw(query, price, uid).Scan(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (crt *cartDatabase) FindProductFromCart(ctx context.Context, cid, pid uint) (bool, error) {

	var exist bool

	query := `SELECT EXISTS (
				SELECT 1
				FROM cart_iteams
				WHERE cart_id = $1
				AND product_id = $2 );`

	err := crt.DB.Raw(query, cid, pid).Scan(&exist).Error

	if err != nil {
		return false, err
	}

	return exist, nil
}

// find cart info
func (crt *cartDatabase) FindCartInfoByCartId(ctx context.Context, cid uint) (domain.CartInfo, error) {

	var body domain.CartInfo

	query := `select * from cart_infos where cart_id = $1;`
	err := crt.DB.Raw(query, cid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (crt *cartDatabase) UpdateCartinfo(ctx context.Context, cid, qty uint, price float64) error {

	var body domain.CartInfo

	// we need to collect the price of the product by product id
	// collect the product price
	// collect the quantity
	// set sub total with quantity * product price

	subtotal := qty * uint(price)

	query := `UPDATE cart_items
				SET sub_total = $1
				WHERE cart_id = $2;`
	err := crt.DB.Raw(query, subtotal, cid).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil

}

func (crt *cartDatabase) CreateCartInfo(ctx context.Context, cid uint) (domain.CartInfo, error) {

	var body domain.CartInfo
	query := `insert into cart_infos (cart_id) values ($1) returning *;`
	err := crt.DB.Raw(query, cid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
func (crt *cartDatabase) Addtocart(ctx context.Context, cart domain.Cart) (domain.Cart, error) {

	var body domain.Cart

	return body, nil

}
