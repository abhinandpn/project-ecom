package repository

import (
	"context"
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
				DB				LIMIT
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

// func (prdt *productDatabase) FindProductInfoByPid(ctx context.Context, pid uint) (domain.ProductInfo, error) {

// 	var body domain.ProductInfo

// 	query := `select * from product_infos where product_id =$1;`

//		err := prdt.DB.Raw(query, pid).Scan(&body).Error
//		if err != nil {
//			return body, err
//		}
//		return body, nil
//	}
func (prdt *productDatabase) UpdateQtyPinfo(ctx context.Context, pid uint, qty uint) error {

	qry := `update product_infos set qty = $1 where product_id = $2 ;`
	err := prdt.DB.Raw(qry, qty, pid).Error
	if err != nil {
		return err
	}
	return nil
}

// ----------------PRODUCT UPDATED---------------

func (p *productDatabase) FindProductByName(name string) (domain.Product, error) {

	var body domain.Product
	query := `select * from products where product_name =$1`
	err := p.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindProductById(id uint) (domain.Product, error) {

	var body domain.Product
	query := `select * from products where id = $1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindProductByBrand(id uint) (domain.Product, error) {

	var body domain.Product
	query := `select * from products where brand_id = $1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindProductByCategory(id uint) (domain.Product, error) {

	var body domain.Product
	query := `select * from products where category_id = $1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindProductBySubCat(id uint) (domain.Product, error) {

	var body domain.Product
	query := `select * from products where sub_category_id = $1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindAllProduct(pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit
	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_infos.size,
				product_images.product_images,
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_infos.product_id
			  ORDER BY
				created_at DESC
			  LIMIT
				$1 OFFSET $2;`

	err := p.DB.Raw(query, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindAllProductWithQuantity(pagination req.PageNation) ([]res.ProductQtyRes, error) {

	var body []res.ProductQtyRes
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				product_infos.colour,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.size,
				product_images.product_images,
				product_infos.quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  ORDER BY
			    product_infos.id ASC
			  LIMIT
				$1 OFFSET $2;`
	err := p.DB.Raw(query, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) CreateProduct(product req.ReqProduct) error {

	tx := p.DB.Begin() // transaction begin
	var Time time.Time
	var ProductTable domain.Product
	var ProductImage domain.ProductImage
	var ProductInfo domain.ProductInfo

	// add in to product table

	query1 := `insert into products(product_name,
					discription,
					brand_id,
					category_id,
					created_at)values ($1,$2,$3,$4,$5) returning id;`
	err := p.DB.Raw(query1, product.ProductName,
		product.Discription,
		product.BrandId,
		product.CategoryID,
		Time).Scan(&ProductTable).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	ProductId := ProductTable.Id
	// add in to productimages table
	query2 := `insert into product_images (product_id,product_images)values ($1,$2) returning id;`
	err = p.DB.Raw(query2, ProductId, product.Image).Scan(&ProductImage).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// ImageId := ProductImage.Id

	// add in to product info table
	query3 := `insert into product_infos (product_id,
					price,
					colour,
					size,
					quantity)values ($1,$2,$3,$4,$5)returning id;`
	err = p.DB.Raw(query3, ProductId,
		product.Price,
		product.Color,
		product.Size, 1).Scan(&ProductInfo).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// commit changes
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (p *productDatabase) DeletProduct(id uint) error {

	tx := p.DB.Begin()

	// delete product info
	query1 := `delete from product_infos where id = $1`
	err := p.DB.Exec(query1, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// delete product image
	// query2 := `delete from product_images where id = $1`
	// err = p.DB.Exec(query2, id).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// delete product table
	query3 := `delete from products where id = $1`
	err = p.DB.Exec(query3, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// commit changes
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (p *productDatabase) FindProductByProductInfo(id uint) (domain.Product, error) {

	var body domain.Product
	query := `select * from product_infos where product_id =$1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) UpdateQuentity(id, qty uint) error {

	query := `update product_infos set quentity =$1 where product_id =$2`
	err := p.DB.Exec(query, qty, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productDatabase) FinBrandByName(name string) (domain.Brand, error) {

	var body domain.Brand
	query := `select * from brands where brand_name =$1`
	err := p.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindCategoryByName(name string) (domain.Category, error) {

	var body domain.Category
	query := `select * from categories where category_name =$1`
	err := p.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) CreateBrand(name, img string) error {

	var body domain.Brand
	query := `insert into brands (brand_name,brand_image)values ($1,$2);`
	err := p.DB.Raw(query, name, img).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productDatabase) DeleteBrand(id uint) error {

	query := `delete from brands where id =$1`
	err := p.DB.Exec(query, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productDatabase) ViewFullBrand() ([]res.ResBrand, error) {

	var body []res.ResBrand
	query := `select * from brands order by id ;`
	err := p.DB.Raw(query).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindBrandByName(name string) (domain.Brand, error) {

	var body domain.Brand
	query := `select * from brands where brand_name = $1`
	err := p.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FIndBrandById(id uint) (domain.Brand, error) {

	var body domain.Brand
	query := ` select * from brands where id =$1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) AddProductImage(id uint, img string) error {

	var body domain.ProductImage
	query := `insert into product_images (product_id,product_images)values ($1,$2);`
	err := p.DB.Raw(query, id, img).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productDatabase) FindImage(img string) (domain.ProductImage, error) {

	var body domain.ProductImage
	query := `select * from product_images where product_images = $1`
	err := p.DB.Raw(query, img).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) FindProductInfoById(id uint) (domain.ProductInfo, error) {

	var body domain.ProductInfo
	query := `select * from product_infos where id = $1;`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) ProductViewByPid(id uint) (res.ResProductOrder, error) {

	// need updation -------------------
	var body res.ResProductOrder
	query := `SELECT 
					
				    p.product_name AS "ProductName",
				    p.discription AS "Description",
				    c.category_name AS "CategoryName",
				    b.brand_name AS "BrandName",
				    pi.size AS "Size",
				    pi.price AS "Price",
				    pi.colour AS "Colour"
				FROM 
				    products p
				JOIN 
				    product_infos pi ON p.id = pi.product_id
				JOIN 
				    brands b ON p.brand_id = b.id
				JOIN 
				    categories c ON p.category_id = c.id
				WHERE 
				    p.id = $1;`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// PRODUCT UPDATION FINAL FUNC

func (p *productDatabase) ProductUpdate(product req.UpdateProduct, id uint) error {

	tx := p.DB.Begin() // transaction begins
	var productTable domain.Product
	var ProductInfo domain.ProductInfo
	var Time time.Time

	// update product info table
	querypf := `update product_infos set price = $1 ,
								colour = $2 ,
								size = $3 ,
								quantity = $4 where id = $5 returning product_id;`
	err := p.DB.Raw(querypf,
		product.Price,
		product.Colour,
		product.Size,
		product.Quantity,
		id).Scan(&ProductInfo).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	// update product table
	queryPr := `update products set product_name = $1,
								discription = $2,
								brand_id = $3,
								category_id = $4,
								updated_at = $5
								where id = $6 returning id ;`
	err = p.DB.Raw(queryPr,
		product.ProductName,
		product.Discription,
		product.BrandId,
		product.CategoryId,
		Time,
		ProductInfo.ProductId).Scan(&productTable).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	// commit changes
	err = tx.Commit().Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// ----------- Sorting -----------

func (p *productDatabase) ListByColour(colour string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  WHERE
				product_infos.colour = $1
			  ORDER BY
			 	 product_infos.id ASC
			  LIMIT
				$2 OFFSET $3;`

	err := p.DB.Raw(query, colour, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) ListBySize(size uint,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  WHERE
				product_infos.size = $1
			  ORDER BY
			  	product_infos.id ASC
			  LIMIT
				$2 OFFSET $3;`

	err := p.DB.Raw(query, size, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) ListByCategory(id uint,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  WHERE
				products.category_id = $1
			  ORDER BY
			  	product_infos.id ASC
			  LIMIT
				$2 OFFSET $3;`

	err := p.DB.Raw(query, id, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productDatabase) ListByBrand(id uint,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  WHERE
				products.brand_id = $1
			  ORDER BY
				product_infos.id ASC
			  LIMIT
				$2 OFFSET $3;`

	err := p.DB.Raw(query, id, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) ListByName(name string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
			  WHERE
				products.product_name = $1
			  ORDER BY
			  	product_infos.id ASC
			  LIMIT
				$2 OFFSET $3;`

	err := p.DB.Raw(query, name, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) ListByPrice(Start, End float64,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
				product_infos.id,
				products.product_name,
				products.discription,
				categories.category_name,
				brands.brand_name,
				product_infos.price,
				product_infos.colour,
				product_images.product_images,
				product_infos.size,
				product_infos,quantity
			  FROM
				products
			  INNER JOIN
				categories ON products.category_id = categories.id
			  INNER JOIN
				brands ON products.brand_id = brands.id
			  INNER JOIN
				product_infos ON products.id = product_infos.product_id
			  INNER JOIN
				product_images ON products.id = product_images.product_id
		      WHERE
			  	product_infos.price BETWEEN $1 AND $2
			  ORDER BY
			 	 product_infos.id ASC
			  LIMIT
				$3 OFFSET $4;`

	err := p.DB.Raw(query, Start, End, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) ListByQuantity(Start, End uint,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT
    product_infos.id,
    products.product_name,
    products.discription,
    categories.category_name,
    brands.brand_name,
    product_infos.price,
    product_infos.colour,
    product_images.product_images,
    product_infos.size,
	product_infos,quantity
		 FROM
		     products
		 INNER JOIN
		     categories ON products.category_id = categories.id
		 INNER JOIN
		     brands ON products.brand_id = brands.id
		 INNER JOIN
		     product_infos ON products.id = product_infos.product_id
		 INNER JOIN
		     product_images ON products.id = product_images.product_id
		 WHERE
		     product_infos.quantity BETWEEN $1 AND $2
			ORDER BY
			  	product_infos.id ASC
		    LIMIT
				$3 OFFSET $4;`
	err := p.DB.Raw(query, Start, End, limit, offset).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) GetProductImage(id uint) (domain.ProductImage, error) {

	var body domain.ProductImage
	query := `select * from product_images where product_id = $1;`
	err := p.DB.Exec(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (p *productDatabase) FindProductPriceByProductInfoId(id uint) (float64, error) {

	var price float64
	query := `select price from product_infos where product_id = $1;`
	err := p.DB.Raw(query, id).Scan(&price).Error
	if err != nil {
		return price, err
	}
	return price, nil
}
