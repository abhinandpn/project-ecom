package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductuseCase interface {
	// Products View
	GetProductInfo(ctx context.Context, ProductId uint) (ProductInfo domain.Product, err error)
	GetProducts(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error)
	// Product Managment
	AddProduct(ctx context.Context, product domain.Product) error

	// Category
	AddCategory(ctx context.Context, category req.CategoryReq) (cat res.CategoryRes, err error)
	Editcategory(ctx context.Context, category domain.Category) error
	DeleteCategory(ctx context.Context, categoryId uint) error
	ViewFullCategory(ctx context.Context, pagination req.PageNation) (category []res.CategoryRes, err error)
}
