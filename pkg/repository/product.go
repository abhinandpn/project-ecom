package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB: db}
}

// todo: PRODUCT FULL FUNCTIONS ( DATABASE RETALED QUER)
/*
* FIND PRODUCT BY ID 	-->
* FIND PRODUCT 			-->
* FIND ALL PRODUCT 		-->
* SAVE PRODUCT 			-->
* UPDATE PRODUCT 		-->
 */

// TODO:
// Find Product
func (pr *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {

	if pr.DB.Raw("SELECT * FROM products WHERE id = ? OR product_name=?", product.Id, product.ProductName).Scan(&product).Error != nil {
		return product, errors.New("faild to get product")
	}
	return product, nil
}

// TODO:
func (pr *productDatabase) FindProductById(ctx context.Context, productId uint) (product domain.Product, err error) {
	query := `Select * from products where id = $1`
	err = pr.DB.Raw(query, productId).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("faild find product with prduct_id %v", productId)
	}
	return product, nil
}

// TODO:
func (pr *productDatabase) FindAllProduct(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	// aliase :: p := product; c := category
	querry := `SELECT
	p.id,
	p.product_name,
	P.discription,
	c.category_name,
	p.price,
	p.discount_price,
	
	pi.colour,
	pi.size,
	pi.brand
  FROM
	products p
  INNER JOIN
	product_infos pi ON p.id = pi.product_id
	
  INNER JOIN
	categories c ON p.category_id = c.id
   ORDER BY created_at DESC LIMIT $1 OFFSET $2
   ;`

	if pr.DB.Raw(querry, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("faild to get products from database")
	}

	return products, nil

}

// TODO:
// Save Product
func (pr *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {

	querry := `INSERT INTO products (product_name, discription,category_id, price, image, created_at) 
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	createdAt := time.Now()

	if pr.DB.Raw(querry, product.ProductName, product.Discription, product.CategoryID,
		product.Price, product.Image, createdAt).Scan(&product).Error != nil {
		// fmt.Errorf(context.Canceled.Error())
		return errors.New("faild to insert product on database")
	}
	query2 := `insert into product_infos(product_id,colour,size,brand)values($1,$2,$3,$4)`

	if pr.DB.Exec(query2, product.Id, product.Info.Colour,
		product.Info.Size, product.Info.Brand).Error != nil {
		return errors.New("faild to insert product_info table on database----------- ")

	}
	return nil
}

// TODO:
func (pr *productDatabase) UpdateProduct(ctx context.Context, product domain.Product) error {

	query := `UPDATE products SET product_name = $1, discription = $2, category_id = $3, 
	price = $4, image = $5, updated_at = $6 WHERE id = $7`

	updatedAt := time.Now()

	if pr.DB.Exec(query, product.ProductName, product.Discription, product.CategoryID,
		product.Price, product.Image, updatedAt, product.Id).Error != nil {
		return errors.New("faild to update product")
	}

	return nil
}

// todo: CATEGORIES
// Categories
func (pr *productDatabase) SaveCategory(ctx context.Context, category domain.Category) error {
	// query := `insert into categories(category_id,)`
	return nil
}
