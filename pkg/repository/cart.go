package repository

import (
	"context"
	"fmt"

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

func (crt *cartDatabase) CreateCart(ctx context.Context, uid uint) error {

	query := `INSERT INTO carts (user_id,total_price) VALUES($1, $2) RETURNING cart_id`

	err := crt.DB.Raw(query, uid, 0).Error
	if err != nil {
		return err
	}
	return nil
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

func (crt *cartDatabase) Addtocart(ctx context.Context, pid uint, uid uint) error {

	// we want user id  and product id
	// get product deail

	// check the user have cart ?
	cart, err := crt.FindCartByUserId(ctx, uid)
	if err != nil {
		return err
	}
	fmt.Println("--------", cart)
	// check the product if exist in the store
	exist, err := crt.FindProductFromCart(ctx, cart.Id, pid)
	if err != nil {
		return err
	}
	fmt.Println("_---------", exist)
	// check the product if exist in cart
	if exist {
		return fmt.Errorf("product alredy exist")
	}
	// if does not exist then add product to the cart items
	// update cart_items
	// va test
	query := `insert into cart_iteams (cart_id,product_id,quantity)values ($1,$2,$3);`
	err = crt.DB.Exec(query, cart.Id, pid, 1).Error
	if err != nil {
		return err
	}
	// update carts it will be in use case

	return nil
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
