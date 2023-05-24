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

// Category
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

func (ct *productUseCase) AddCategory(ctx context.Context, category domain.Category) (domain.Category, error) {

	// find if exist or not
	body, err := ct.productRepo.FindCategoryById(ctx, category.Id)
	if err != nil {
		fmt.Println("xtxtxtxtxtxtxtxtxt")
		return category, err
	} else if body.Id == 0 {
		return category, fmt.Errorf("invalid category_id %v", body.Id)
	}
	// add category
	err = ct.productRepo.CreateCategory(ctx, body)
	if err != nil {
		return body, err
	}
	// responce
	log.Printf("successfully product saved\n\n")
	return body, nil
}

func (ct *productUseCase) UpdateCategory(ctx context.Context, category domain.Category) error {

	// chek the category
	body, err := ct.productRepo.FindCategoryById(ctx, category.Id)
	if err != nil {
		return err
	} else if body.Id == 0 {
		return fmt.Errorf("invalid product_id %v", body.Id)
	}
	// update
	_, err = ct.productRepo.UpdateCategory(ctx, category)
	if err != nil {
		return err
	}
	//responce
	return nil
}

func (ct *productUseCase) DeleteCategory(ctx context.Context, id uint) error {

	// check if exist or not
	body, err := ct.productRepo.FindCategoryById(ctx, id)
	if err != nil {
		return err
	}
	// delete
	err = ct.DeleteCategory(ctx, body.Id)
	if err != nil {
		return nil
	}
	// response
	return nil
}

func (pr *productUseCase) ViewFullCategory(ctx context.Context, pagination req.PageNation) ([]res.CategoryRes, error) {

	return pr.productRepo.FindAllCategory(ctx, pagination)

}
