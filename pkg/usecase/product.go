package usecase

import (
	"context"
	"fmt"
	"log"

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
	if cat.Id != 0 {
		err = fmt.Errorf("category name : %v", cat.CategoryName)
		return err
	} else if cat.Id == 0 {
		err = fmt.Errorf("category id does not exist create first ")
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

func (pr *productUseCase) UpdateProduct(ctx context.Context, product domain.Product) error {

	// check the product exist or not
	Pid := product.Id
	body, err := pr.productRepo.FindProductById(ctx, Pid)
	if err != nil {
		return err
	} else if body.Id == 0 {
		return fmt.Errorf("invalid product_id %v", body.Id)
	}

	// if exist update
	_, err = pr.productRepo.UpdateProduct(ctx, body)
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
	} else if body.Id == 0 {
		return fmt.Errorf("invalid product_id %v", body.Id)
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
