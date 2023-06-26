package usecase

import (
	"context"
	"errors"
	"fmt"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
	"github.com/jinzhu/copier"

	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

// to get a new instance of productUseCase
func NewProductUseCase(ProductRepo interfaces.ProductRepository) service.ProductuseCase {
	return &productUseCase{productRepo: ProductRepo}
}

// Product
/*
func (pr *productUseCase) FindProductByName(ctx context.Context, name string) (domain.Product, error) {

	body, err := pr.productRepo.FindProductByName(ctx, name)
	if err != nil {
		return body, err
	}
	return body, nil
}

// -------------------AddProduct-------------------

func (pr *productUseCase) AddProduct(ctx context.Context, product req.ReqProduct) error {

	// check Given product is exist or not
	name := product.ProductName

	body, err := pr.FindProductByName(ctx, name)
	if err != nil {
		return err
	}

	if body.Id != 0 {
		err = fmt.Errorf("product alredy exist name : %v", body.ProductName)
		return err
	}

	// check the category if exist or not
	ct, err := pr.FindCategoryById(ctx, product.CategoryID)
	if err != nil {
		return err
	}

	cat, err := pr.FindCategoryByname(ctx, ct.CategoryName)
	if err != nil {
		return err
	}
	// cat res
	if cat.Id == 0 {
		err = fmt.Errorf("category name : %v", cat.CategoryName)
		return err
	}

	// product res
	if body.Id != 0 {
		err = fmt.Errorf("product alredy exist with this name : %v", name)
		return err
	}

	// if its found then create a new product with new data
	err = pr.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return err
	}

	log.Printf("successfully product saved\n\n")
	return nil
}

func (pr *productUseCase) UpdateProduct(ctx context.Context, product req.ReqProduct, id uint) error {

	// check if the product exist or not
	body, err := pr.productRepo.FindProductById(ctx, id)
	if err != nil {
		return err
	}
	if body.Id == 0 {
		err = fmt.Errorf("unabel to find product by Id %v", id)
		return err
	}

	// check if exist in category
	category, err := pr.productRepo.FindCategoryById(ctx, product.CategoryID)
	if err != nil {
		return err
	}
	if category.Id == 0 {
		err = fmt.Errorf("unabel to find category by Id %v", id)
		return err
	}
	// if exist update product table
	err = pr.productRepo.UpdateProduct(ctx, product, id)
	if err != nil {
		return err
	}
	// responce
	return nil
}

func (pr *productUseCase) DeleteProduct(ctx context.Context, id uint) error {

	// check the product exist or not
	body, err := pr.productRepo.FindProductById(ctx, id)
	if err != nil {
		return err
	}
	if body.Id == 0 {
		return fmt.Errorf("invalid product id %v", body.Id)
	}

	// if exist then delete
	err = pr.productRepo.DeletProduct(ctx, body.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productUseCase) FindProductById(ctx context.Context, id uint) (res.ProductResponce, error) {

	var Product res.ProductResponce
	body, err := pr.productRepo.FindProductById(ctx, id)
	if err != nil {
		return Product, err
	}

	err = copier.Copy(&Product, &body)
	if err != nil {
		return Product, err
	}
	return Product, nil
}

func (pr *productUseCase) ViewFullProduct(ctx context.Context, pagination req.PageNation) ([]res.ProductResponce, error) {

	body, err := pr.productRepo.ViewFullProduct(ctx, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}
*/
// CATEGORY USECASE

// -------------------FindCategoryById-------------------

func (ct *productUseCase) FindCategoryById(ctx context.Context, id uint) (res.CategoryRes, error) {

	var category res.CategoryRes
	body, err := ct.productRepo.FindCategoryById(ctx, id)
	if err != nil {
		return category, err
	}

	err = copier.Copy(&category, &body)
	if err != nil {
		return category, err
	}
	return category, nil
}

// -------------------FindCategoryByName-------------------

func (ct *productUseCase) FindCategoryByname(ctx context.Context, name string) (domain.Category, error) {

	body, err := ct.productRepo.FindCategoryByname(ctx, name)
	if err != nil {
		return body, err
	}
	return body, nil

}

// -------------------AddCategory-------------------

func (ct *productUseCase) AddCategory(ctx context.Context, name string) (domain.Category, error) {

	// create response
	var body domain.Category

	// find with name if exist
	exist, err := ct.FindCategoryByname(ctx, name)
	if err != nil {
		return exist, err
	}

	// if exist return
	if exist.Id != 0 {
		res := fmt.Errorf("category alredy exist with this name :%v ", name)
		return exist, res
	}
	// or create
	body, err = ct.productRepo.CreateCategory(ctx, name)
	if err != nil {
		return body, err
	}

	// return
	return body, nil

}

// -------------------Updatecategory-------------------

func (ct *productUseCase) UpdateCategory(ctx context.Context, category req.UpdateCategoryReq) error {

	// check if the category exist or not
	body, err := ct.productRepo.FindCategoryByname(ctx, category.OldCategory)
	if err != nil {
		return err
	}

	if body.Id == 0 {
		res := fmt.Errorf("category not found with this name :  %v", category.OldCategory)
		return res
	}
	// update
	body, err = ct.productRepo.UpdateCategory(ctx, category)
	if err == nil {
		return err
	}

	// response
	return nil
}

// -------------------Deletecategory-------------------

func (ct *productUseCase) DeleteCategory(ctx context.Context, name string) error {

	// check if exist or not
	body, err := ct.productRepo.FindCategoryByname(ctx, name)
	if err != nil {
		return err
	}
	// check
	if body.Id == 0 {
		res := fmt.Errorf("category not found with this name :  %v", name)
		return res
	}
	// delete
	_, err = ct.productRepo.DeleteCategory(ctx, name)
	if err != nil {
		return nil
	}
	// response
	return nil
}

// -------------------ViewFullcategory-------------------

func (pr *productUseCase) ViewFullCategory(ctx context.Context,
	pagination req.PageNation) ([]res.CategoryRes, error) {

	return pr.productRepo.FindAllCategory(ctx, pagination)

}

// ------------------------SUB CATEGORY------------------------

func (sub *productUseCase) AddSubcategory(ctx context.Context, cid uint, name string) (domain.SubCategory, error) {

	body, err := sub.productRepo.CreateSubCategory(ctx, cid, name)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) SubNameUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error) {

	body, err := sub.productRepo.ChangeSubCatName(ctx, id, name)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) SubCatUpdate(ctx context.Context, id uint, name string) (domain.SubCategory, error) {

	body, err := sub.productRepo.ChangeSubCatCatogeryName(ctx, id, name)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) DeleteSubCate(ctx context.Context, name string) error {

	err := sub.productRepo.DeleteSubCategory(ctx, name)
	return err
}
func (sub *productUseCase) ListALlSubCate(ctx context.Context, pagination req.PageNation) ([]res.SubCategoryRes, error) {

	body, err := sub.productRepo.ListllSubCategory(ctx, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) FindSubById(ctx context.Context, id uint) (domain.SubCategory, error) {

	body, err := sub.productRepo.FindSubCtById(ctx, id)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) FindSubByName(ctx context.Context, name string) (domain.SubCategory, error) {

	body, err := sub.productRepo.FindSubCtByName(ctx, name)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) FindSubByCatId(ctx context.Context, id uint) (domain.SubCategory, error) {

	body, err := sub.productRepo.FindSubCtByCategoryId(ctx, id)
	if err != nil {
		return body, err
	}
	return body, nil
}
func (sub *productUseCase) FindSubByCatName(ctx context.Context, name string) (domain.SubCategory, error) {

	body, err := sub.productRepo.FindSubCtByCategoryName(ctx, name)
	if err != nil {
		return body, err
	}
	return body, nil
}

// func (sub *productUseCase) FindSubByCatName(ctx context.Context, name string) (domain.SubCategory, error) {

// 	body, err := sub.productRepo.FindSubCtByCategoryName(ctx, name)
// 	if err != nil {
// 		return body, err
// 	}
// 	return body, nil
// }

// -------------PRODUCT UPDATED-------------

func (p *productUseCase) GetProductByName(name string) (domain.Product, error) {

	product, err := p.productRepo.FindProductByName(name)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productUseCase) GetProductById(id uint) (domain.Product, error) {

	product, err := p.productRepo.FindProductById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productUseCase) GetAllProducts(pagination req.PageNation) ([]res.ProductResponce, error) {

	products, err := p.productRepo.FindAllProduct(pagination)
	if err != nil {
		return products, err
	}
	return products, nil
}

func (p *productUseCase) ListproductByBrand(brand string) (domain.Product, error) {

	var product domain.Product
	// find brand id
	brands, err := p.productRepo.FinBrandByName(brand)
	if err != nil {
		return product, err
	}

	// list product
	product, err = p.productRepo.FindProductByBrand(brands.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productUseCase) ListProductByCategory(category string) (domain.Product, error) {

	var product domain.Product
	// find category
	body, err := p.productRepo.FindCategoryByName(category)
	if err != nil {
		return product, err
	}

	// list product
	product, err = p.productRepo.FindProductByCategory(body.Id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productUseCase) CreateProduct(product req.ReqProduct) error {

	var ctx context.Context // contexte
	// check the name exist
	Prdt, err := p.productRepo.FindProductByName(product.ProductName)
	if err != nil {
		return err
	}
	if Prdt.Id == 0 {

		// find category
		Cat, err := p.productRepo.FindCategoryById(ctx, product.CategoryID)

		// if does not exist return
		if err != nil {
			return err
		}
		// find brand

		// if exist
		if Cat.Id != 0 {
			brand, err := p.productRepo.FIndBrandById(product.BrandId)
			if err != nil {
				return err
			}
			if brand.Id != 0 {
				err = p.productRepo.CreateProduct(product)
				if err != nil {
					return err
				}
			} else {
				res := errors.New("brand does not exist")
				return res
			}

		} else {
			res := errors.New("category does not exist")
			return res
		}

	} else {
		res := errors.New("product alredy exist")
		return res
	}
	// return
	return nil
}

func (p *productUseCase) DeleteProduct(id uint) error {

	// find product if exist or not
	product, err := p.productRepo.FindProductById(id)
	if err != nil {
		return err
	}

	// if exist delet
	if product.Id != 0 {

		// delete product
		err := p.productRepo.DeletProduct(id)
		if err != nil {
			return err
		}
	}

	// response
	return nil
}

func (p *productUseCase) UpdateProduct(product res.ResProduct, id uint) error {

	// find product
	prdt, err := p.productRepo.FindProductByName(product.ProductName)
	if err != nil {
		return err
	}
	if prdt.Id == 0 {
		return errors.New("product does not exist")
	}
	err = p.productRepo.UpdateProduct(product, prdt.Id)
	if err != nil {
		return err
	}
	return err
}

func (p *productUseCase) UpdateQuentity(id, qty uint) error {

	// find product
	product, err := p.productRepo.FindProductById(id)
	if err != nil {
		return err
	}
	if product.Id == 0 {
		return errors.New("product does not exist")
	}
	// update
	err = p.productRepo.UpdateQuentity(id, qty)
	if err != nil {
		return err
	}
	// response
	return err
}

func (p *productUseCase) CreateBrand(name, img string) error {

	// check if exist or not
	brand, err := p.productRepo.FinBrandByName(name)
	if err != nil {
		return err
	}

	// if does not exist create
	if brand.Id != 0 {
		if brand.Id != 0 {
			res := fmt.Errorf("category alredy exist with this name :%v ", name)
			return res
		}
	} else {
		err = p.productRepo.CreateBrand(name, img)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *productUseCase) DeleteBrand(id uint) error {

	// finc if id exit
	brand, err := p.productRepo.FIndBrandById(id)
	if err != nil {
		return err
	}
	// if exist delete
	if brand.Id != 0 {
		err = p.productRepo.DeleteBrand(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *productUseCase) ViewFullBrand() (res.ResBrand, error) {

	body, err := p.productRepo.ViewFullBrand()
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productUseCase) AddImage(id uint, name string) error {

	// check value
	image, err := p.productRepo.FindImage(name)
	if err != nil {
		return err
	}
	if image.Id != 0 {
		res := errors.New("image alredy exist")
		return res
	}
	// add image
	err = p.productRepo.AddProductImage(id, name)
	if err != nil {
		return err
	}
	return nil
}

// ------ SORTING -------

func (p *productUseCase) GetByColour(colour string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	body, err := p.productRepo.ListByColour(colour, pagination)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (p *productUseCase) GetBySize(size int,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	s := uint(size)
	body, err := p.productRepo.ListBySize(s, pagination)
	if err != nil {
		return body, err
	}

	return body, nil
}

func (p *productUseCase) GetByCategory(name string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var ctx context.Context
	var body []res.ProductResponce

	category, err := p.productRepo.FindCategoryByname(ctx, name)
	if err != nil {
		return body, err
	}

	body, err = p.productRepo.ListByCategory(category.Id, pagination)
	if err != nil {
		return body, err
	}

	return body, nil

}

func (p *productUseCase) GetByBrand(name string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	var body []res.ProductResponce

	brand, err := p.productRepo.FindBrandByName(name)
	if err != nil {
		return body, err
	}

	body, err = p.productRepo.ListByBrand(brand.Id, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productUseCase) GetByName(name string,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	body, err := p.productRepo.ListByName(name, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productUseCase) GetByPrice(Start, End int,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	st := float64(Start)
	ed := float64(End)

	body, err := p.productRepo.ListByPrice(st, ed, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *productUseCase) GetByQuantity(Start, End int,
	pagination req.PageNation) ([]res.ProductResponce, error) {

	st := uint(Start)
	ed := uint(End)

	body, err := p.productRepo.ListByQuantity(st, ed, pagination)
	if err != nil {
		return body, err
	}
	return body, nil
}
