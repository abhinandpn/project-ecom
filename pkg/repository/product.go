package repository

import (
	"context"
	"errors"
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

// Find Product
func (pr *productDatabase) FindProduct(ctx context.Context, product domain.Product) (domain.Product, error) {

	if pr.DB.Raw("SELECT * FROM products WHERE id = ? OR product_name=?", product.Id, product.ProductName).Scan(&product).Error != nil {
		return product, errors.New("faild to get product")
	}
	return product, nil
}
func (pr *productDatabase) FindAllProduct(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	// aliase :: p := product; c := category
	querry := `SELECT pi.id,p.product_name,p.discription,c.category_name,p.price,
	p.discount_price,p.image,pi.colour,pi.size,
	pi.brand FROM product_infos pi JOIN products p ON pi.product_id = p.id JOIN categories c ON 
	p.category_id = c.id ORDER BY created_at DESC LIMIT $1 OFFSET $2;`

	if pr.DB.Raw(querry, limit, offset).Scan(&products).Error != nil {
		return products, errors.New("faild to get products from database")
	}

	return products, nil

}

// Save Product
func (pr *productDatabase) SaveProduct(ctx context.Context, product domain.Product) error {

	querry := `INSERT INTO products (product_name, description, category_id, price, image, created_at) 
	VALUES($1, $2, $3, $4, $5, $6)`

	createdAt := time.Now()
	if pr.DB.Exec(querry, product.ProductName, product.Discription, product.CategoryId,
		product.Price, product.Image, createdAt).Error != nil {
		return errors.New("faild to insert product on database")
	}
	return nil
}
