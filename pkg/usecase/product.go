package usecase

import (
	"context"
	"fmt"
	"log"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"

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
