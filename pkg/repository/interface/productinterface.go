package interfaces

import (
	"context"

	"github.com/abhinandpn/project-ecom/pkg/util/req"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
)

type ProductRepository interface {
	// List all product
	FindAllProduct(ctx context.Context, pagination req.PageNation) (products []res.ProductResponce, err error)
}
