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

// Product

func (pr *productDatabase) FindProductById(ctx context.Context, PId uint) (domain.Product, error) {

	var product domain.Product
	query := `Select * from products where id = $1`

	err := pr.DB.Raw(query, PId).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("faild find To product with prduct_id : %v", PId)
	}
	return product, nil
}

func (pr *productDatabase) ViewFullProduct(ctx context.Context, pagination req.PageNation) (product []res.ProductResponce, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit
	var ProductTable res.ProductResponce
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
   ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	err = pr.DB.Raw(querry, limit, offset).Scan(&ProductTable).Error
	if err != nil {
		return product, errors.New("faild to get product details from database")
	}
	return product, nil
}

func (pr *productDatabase) CreateProduct(ctx context.Context, product domain.Product) error {

	// verify the product by id
	product, err := pr.FindProductById(ctx, product.Id)
	if err != nil {
		return err
	}

	querry := `INSERT INTO products (product_name, 
					discription,
					category_id, 
					price, image, 
					created_at) 
				VALUES($1, $2, $3, $4, $5, $6) RETURNING id`

	createdAt := time.Now()
	var ResProduct res.ProductResponce

	if pr.DB.Raw(querry,
		product.ProductName,
		product.Discription,
		product.CategoryID,
		product.Price,
		product.Image,
		createdAt).Scan(&ResProduct).Error != nil {
		return errors.New("faild to insert product on database")
	}

	query2 := `insert into product_infos(product_id,
		colour,
		size,
		brand)values($1,$2,$3,$4)`

	if pr.DB.Raw(query2,
		product.Id,
		product.Info.Colour,
		product.Info.Size,
		product.Info.Brand).Scan(ResProduct).Error != nil {
		return errors.New("faild to insert product_info table on database----------- ")

	}
	return nil
}

func (pr *productDatabase) DeletProduct(ctx context.Context, PId uint) error {

	// verify the product by id
	product, err := pr.FindProductById(ctx, PId)
	if err != nil {
		return err
	}

	// delete the product_info
	query1 := `delete from product_infos where product_id = $1;`
	err = pr.DB.Raw(query1).Error
	if err != nil {
		return errors.New("failed to delete from product_infos table")
	}

	// Delete from Products
	query2 := `delete from products where id = $1;`
	err = pr.DB.Raw(query2, product.Id).Error
	if err != nil {
		return errors.New("failed to delete from products table")
	}

	return nil
}

func (pr *productDatabase) UpdateProduct(ctx context.Context, info domain.Product) (domain.Product, error) {

	query := `UPDATE products SET 
					product_name = $1, 
					discription = $2, 
					category_id = $3, 
					price = $4, 
					image = $5, 
					updated_at = $6 
				WHERE id = $7`

	updatedAt := time.Now()
	var product domain.Product

	err := pr.DB.Exec(query, info.ProductName,
		info.Discription,
		info.CategoryID,
		info.Price,
		info.Image,
		updatedAt,
		info.Id).Scan(product).Error

	if err != nil {
		return info, nil
	}

	return product, nil
}

// Categories New updated

func (ct *productDatabase) FindCategoryById(ctx context.Context, CId uint) (domain.Category, error) {

	var category domain.Category
	query := `select * from categories where id = $1;`

	err := ct.DB.Raw(query, CId).Scan(&category).Error
	if err != nil {
		return category, fmt.Errorf("faild to find to product with this category id %v", CId)
	}

	return category, nil
}

func (ct *productDatabase) CreateCategory(ctx context.Context, Category domain.Category) error {

	// verify the product by id
	body, err := ct.FindCategoryById(ctx, Category.Id)
	if err != nil {
		return err
	}
	fmt.Println("body body body ------ ??? > ", body)
	// var body domain.Category
	// if its not exist then create new one using this fileds
	query := `insert into categories (id,category_name)values ($1,$2);`
	fmt.Println("xxxxxxxxxxx-------- > id ----- >", Category.Id)

	fmt.Println("xxxxxxxxxxx-------- > name ------- >", Category.CategoryName)

	err = ct.DB.Raw(query, body.Id, body.CategoryName).Scan(body).Error
	if err != nil {
		return err
	}
	// return
	return nil
}

func (ct *productDatabase) DeleteCategory(ctx context.Context, id uint) error {

	// find the category
	category, err := ct.FindCategoryById(ctx, id)
	if err != nil {
		return err
	}
	// if we get delete
	query := `delete from categories where id = $1;`
	err = ct.DB.Raw(query, category.Id).Error
	if err != nil {
		return err
	}
	// retun
	return nil
}

func (ct *productDatabase) UpdateCategory(ctx context.Context, info domain.Category) (domain.Category, error) {

	var body domain.Category
	// find the category
	category, err := ct.FindCategoryById(ctx, info.Id)
	if err != nil {
		return category, err
	}
	// update
	query := `update categories set category_name = $1 where id = $2;`
	err = ct.DB.Raw(query, info.CategoryName, info.Id).Scan(&body).Error
	if err != nil {
		return category, err
	}
	// return
	return body, nil

}

func (pr *productDatabase) FindAllCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	var category []res.CategoryRes

	query := `select * from categories order by id asc limit $1 offset $2;`
	err := pr.DB.Raw(query, limit, offset).Scan(&category).Error

	if err != nil {
		return category, fmt.Errorf("failed to get categories from database")
	}

	return category, nil
}
