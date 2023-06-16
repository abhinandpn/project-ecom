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

// -------------------FindProductById-------------------

func (pr *productDatabase) FindProductById(ctx context.Context, PId uint) (domain.Product, error) {

	var product domain.Product
	query := `Select * from products where id = $1`

	err := pr.DB.Raw(query, PId).Scan(&product).Error
	if err != nil {
		return product, fmt.Errorf("faild find To product with prduct_id : %v", PId)
	}
	return product, nil
}

func (pr *productDatabase) FindProductByName(ctx context.Context, name string) (domain.Product, error) {

	var body domain.Product

	query := `select * from products where product_name = ?;`

	err := pr.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// -------------------ViewFullProduct-------------------

func (pr *productDatabase) ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit
	var product []res.ProductResponce
	// var ProductTable res.ProductResponce
	// aliase :: p := product; c := category
	querry := `SELECT
					p.id,
					p.product_name,
					p.discription,
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
				  ORDER BY
					created_at DESC
				  LIMIT
					$1 OFFSET $2;`

	err := pr.DB.Raw(querry, limit, offset).Scan(&product).Error
	if err != nil {
		return product, errors.New("faild to get product details from database")
	}
	return product, nil
}

// -------------------CreateProduct-------------------

func (pr *productDatabase) CreateProduct(ctx context.Context, product req.ReqProduct) error {

	createdAt := time.Now()

	var prdt domain.Product
	var prdtinfo domain.ProductInfo

	querry := `INSERT INTO products (product_name, 
					discription,
					category_id, 
					price, image, 
					created_at) 
				VALUES($1, $2, $3, $4, $5, $6) returning id`

	// var ResProduct res.ProductResponce

	if pr.DB.Raw(querry,
		product.ProductName,
		product.Discription,
		product.CategoryID,
		product.Price,
		product.Image,
		createdAt).Scan(&prdt).Error != nil {
		return errors.New("faild to insert product on database")
	}

	query2 := `insert into product_infos(product_id,
		colour,
		size,
		brand)values($1,$2,$3,$4)`

	productId := prdt.Id

	if pr.DB.Raw(query2,
		productId,
		product.Color,
		product.Size,
		product.Brand).Scan(&prdtinfo).Error != nil {
		return errors.New("faild to insert product_info table on database----------- ")

	}
	return nil
}

// -------------------DeleteProduct-------------------

func (pr *productDatabase) DeletProduct(ctx context.Context, PId uint) error {

	var prdtinfo domain.ProductInfo
	var prdt domain.Product

	// delete the product_info
	query1 := `delete from product_infos where product_id = $1 ;`
	err := pr.DB.Raw(query1, PId).Scan(&prdtinfo).Error
	if err != nil {
		return errors.New("failed to delete from product_infos table")
	}

	// Delete from Products
	query2 := `delete from products where id = $1;`
	err = pr.DB.Raw(query2, PId).Scan(&prdt).Error
	if err != nil {
		return errors.New("failed to delete from products table")
	}

	return nil
}

// -------------------UpdateProduct-------------------

func (pr *productDatabase) UpdateProduct(ctx context.Context, info req.ReqProduct, id uint) error {

	query1 := `UPDATE products SET 
					product_name = $1, 
					discription = $2, 
					category_id = $3, 
					price = $4, 
					image = $5,
					update_at = $6 
				WHERE id = $7 returning id;`

	updatedAt := time.Now()

	var product domain.Product

	err := pr.DB.Raw(query1, info.ProductName,
		info.Discription,
		info.CategoryID,
		info.Price,
		info.Image,
		updatedAt, id).Scan(&product).Error

	if err != nil {
		return err
	}

	query2 := `update product_infos set colour = $1,
					size = $2,
					brand = $3;
					where product_id = $4,`

	err = pr.DB.Raw(query2, info.Color, info.Size, info.Brand, id).Error

	if err != nil {
		return err
	}
	return nil
}

// Categories New updated

// -------------------FindcategoryById-------------------

func (ct *productDatabase) FindCategoryById(ctx context.Context, CId uint) (domain.Category, error) {

	var category domain.Category
	query := `select * from categories where id = $1;`

	err := ct.DB.Raw(query, CId).Scan(&category).Error
	if err != nil {
		return category, fmt.Errorf("faild to find to product with this category id %v", CId)
	}

	return category, nil
}

// -------------------FindcategoryByName-------------------

func (ct *productDatabase) FindCategoryByname(ctx context.Context, name string) (domain.Category, error) {

	var category domain.Category

	query := `select * from categories where category_name = ?`

	err := ct.DB.Raw(query, name).Scan(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil

}

// -------------------CreateCategory-------------------

func (ct *productDatabase) CreateCategory(ctx context.Context, name string) (domain.Category, error) {

	var body domain.Category

	query := `insert into categories (category_name)values ($1);`

	err := ct.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// -------------------FindFullCategory-------------------

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

// -------------------UpdateCategory-------------------

func (ct *productDatabase) UpdateCategory(ctx context.Context, body req.UpdateCategoryReq) (domain.Category, error) {

	var category domain.Category

	query := `update categories set category_name = $1 where category_name = $2;`

	err := ct.DB.Raw(query, body.Newcategory, body.OldCategory).Scan(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

// -------------------DeleteCategory-------------------

func (ct *productDatabase) DeleteCategory(ctx context.Context, name string) (domain.Category, error) {

	var body domain.Category
	// if we get delete
	query := `delete from categories where category_name = $1;`

	err := ct.DB.Raw(query, name).Scan(&body).Error

	if err != nil {
		return body, err
	}

	// retun
	return body, nil
}

// -------------------------SUB-CATEGORY-------------------------
// Find by name
func (subct *productDatabase) FindSubCtByName(ctx context.Context, name string) (domain.SubCategory, error) {

	var body domain.SubCategory
	query := `select * from sub_categories where category_name =?;`

	err := subct.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// find by id
func (subct *productDatabase) FindSubCtById(ctx context.Context, id uint) (domain.SubCategory, error) {

	var body domain.SubCategory
	query := `select * from sub_categories where id =?`

	err := subct.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// find by sub category name

func (subct *productDatabase) FindSubCtByCategoryName(ctx context.Context, name string) (domain.SubCategory, error) {

	var body domain.SubCategory

	// find category with name
	cat, err := subct.FindCategoryByname(ctx, name)
	if err != nil {
		return body, err
	}
	// find with this catogery id
	query := `select * from sub_categories where category_id =?;`
	err = subct.DB.Raw(query, cat.Id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// find sub category by id
func (subct *productDatabase) FindSubCtByCategoryId(ctx context.Context, id uint) (domain.SubCategory, error) {

	var body domain.SubCategory

	// find category with name
	cat, err := subct.FindCategoryById(ctx, id)
	if err != nil {
		return body, err
	}
	query := `select * from sub_categories where category_id =?;`
	err = subct.DB.Raw(query, cat.Id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// list all sub category
func (subct *productDatabase) ListllSubCategory(ctx context.Context, pagination req.PageNation) ([]res.SubCategoryRes, error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit
	var body []res.SubCategoryRes

	query := `SELECT c.id, c.category_name, sc.category_id, sc.sub_category_name
					FROM categories c
				INNER JOIN sub_categories sc
					ON c.id = sc.category_id
				ORDER BY
					id DESC
				LIMIT
					$1 OFFSET $2;`

	err := subct.DB.Raw(query, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// create sub category

func (subct *productDatabase) CreateSubCategory(ctx context.Context, cid uint, name string) (domain.SubCategory, error) {

	var body domain.SubCategory

	// findiing category
	category, err := subct.FindCategoryById(ctx, cid)
	if err != nil {
		return body, err
	}
	// craeting new sub category
	query := `insert into sub_categories (category_id,sub_category_name)values ($1,$2);`
	err = subct.DB.Raw(query, category.Id, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// deleting sub category

func (subct *productDatabase) DeleteSubCategory(ctx context.Context, name string) error {

	query := `delete from sub_categories  where sub_category_name =$1`
	err := subct.DB.Exec(query).Error
	if err != nil {
		return err
	}
	return nil
}

// changing sub cat name
func (subct *productDatabase) ChangeSubCatName(ctx context.Context, id uint, name string) (domain.SubCategory, error) {

	var body domain.SubCategory

	// find the sub category details
	subcategory, err := subct.FindSubCtById(ctx, id)
	if err != nil {
		return body, err
	}
	// find the
	query := `UPDATE sub_categories
			  SET sub_category_name = $1
			  WHERE sub_category_name = $2;`
	err = subct.DB.Raw(query, name, subcategory.SubCategoryName).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// change category for a sub category
func (subct *productDatabase) ChangeSubCatCatogeryName(ctx context.Context, id uint, name string) (domain.SubCategory, error) {

	var body domain.SubCategory
	// find the sub category
	subcategory, err := subct.FindSubCtByCategoryName(ctx, name)
	if err != nil {
		return body, err
	}
	query := `UPDATE sub_categories
				SET category_id = $1
				WHERE sub_category_name = $2;`
	err = subct.DB.Raw(query, id, subcategory.SubCategoryName).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
func (prdt *productDatabase) FindProductInfoByPid(ctx context.Context, pid uint) (domain.ProductInfo, error) {

	var body domain.ProductInfo

	query := `select * from product_infos where product_id =$1;`

	err := prdt.DB.Raw(query, pid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
func (prdt *productDatabase) UpdateQtyPinfo(ctx context.Context, pid uint, qty uint) error {

	qry := `update product_infos set qty = $1 where product_id = $2 ;`
	err := prdt.DB.Raw(qry, qty, pid).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productDatabase) FindProductByPrinfo(pfid uint) (uint, error) {

	var prinfo domain.ProductInfo
	var id uint
	qry := `select * from product_infos where id = $1;`
	err := p.DB.Raw(qry, pfid).Scan(&prinfo).Error
	id = prinfo.Id
	// fmt.Println(prinfo)
	if err != nil {
		return id, err
	}
	return id, nil
}
