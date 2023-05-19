package usecase

import (
	"context"
	"errors"
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
func (pr *productUseCase) AddProduct(ctx context.Context, product domain.Product) error {

	// check Given product is exist or not
	Pid := product.Id
	body, err := pr.productRepo.FindProductById(ctx, Pid)
	if err != nil {
		return err
	} else if body.Id == 0 {
		return fmt.Errorf("invalid product_id %v", body.Id)
	}

	// if its found then create a new product with new data
	err = pr.productRepo.CreateProduct(ctx, body)
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

//
//
//
//

// Category
func (pr *productUseCase) AddCategory(ctx context.Context, category req.CategoryReq) (res.CategoryRes, error) {

	exitstingCategory, _ := pr.productRepo.FindCategoryById(ctx, category.Id)

	if exitstingCategory.Id == 0 {
		return res.CategoryRes{}, errors.New("category already exists")
	}

	err := pr.productRepo.SaveCategory(ctx, category)
	return res.CategoryRes(category), err
}

func (pr *productUseCase) Editcategory(ctx context.Context, category domain.Category) error {

	// check if the category exist or not

	var checkCategory = domain.Category{
		CategoryID: category.Id}

	if checkCategory, err := pr.productRepo.FindCategoryById(ctx, checkCategory.CategoryID); err != nil {
		return err
	} else if checkCategory.Id == 0 {
		return err
	}

	// if we get then update
	var newCate domain.Category
	copier.Copy(&newCate, &category)

	err := pr.productRepo.UpdateCatrgoryName(ctx, newCate)
	if err != nil {
		return err
	}

	return nil
}

func (pr *productUseCase) DeleteCategory(ctx context.Context, categoryId uint) error {

	// check if the category exist or not

	var checkCategory = domain.Category{
		CategoryID: categoryId}

	if checkCategory, err := pr.productRepo.FindCategoryById(ctx, checkCategory.CategoryID); err != nil {
		return err
	} else if checkCategory.Id == 0 {
		return err
	}

	var body domain.Category
	copier.Copy(&body, &categoryId)

	err := pr.productRepo.DeletCategory(ctx, body)
	if err != nil {
		return err
	}

	return nil
}

func (pr *productUseCase) ViewFullCategory(ctx context.Context, pagination req.PageNation) (category []res.CategoryRes, err error) {

	return pr.productRepo.FindAllCategory(ctx, pagination)

}
