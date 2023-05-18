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

func (pr *productUseCase) GetProducts(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error) {
	return pr.productRepo.FindAllProduct(ctx, pagination)
}

func (pr *productUseCase) AddProduct(ctx context.Context, product domain.Product) error {

	// check Given product is exist or not
	if product, err := pr.productRepo.FindProduct(ctx, product); err != nil {
		return err
	} else if product.Id != 0 {
		return fmt.Errorf("product already exist with %s product name", product.ProductName)
	}
	log.Printf("successfully product saved\n\n")

	return pr.productRepo.SaveProduct(ctx, product)
}

func (pr *productUseCase) GetProductInfo(ctx context.Context, ProductId uint) (ProductInfo domain.Product, err error) {
	return ProductInfo, nil
}

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
