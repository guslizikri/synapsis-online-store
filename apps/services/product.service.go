package services

import (
	"context"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
)

type RepoProductIF interface {
	CreateProduct(ctx context.Context, model entity.Product) (err error)
	GetAllProduct(ctx context.Context, model entity.ProductQuery) (products []entity.Product, err error)
	// GetProductBySKU(ctx context.Context, sku string) (product Product, err error)
}

type ServiceProduct struct {
	repo RepoProductIF
}

func NewServiceProduct(repo RepoProductIF) *ServiceProduct {
	return &ServiceProduct{
		repo: repo,
	}
}

func (s ServiceProduct) CreateProduct(ctx context.Context, req request.CreateProductRequestPayload) (err error) {
	productEntity := entity.NewProductFromCreateProductRequest(req)

	err = productEntity.Validate()
	if err != nil {
		return
	}

	err = s.repo.CreateProduct(ctx, productEntity)
	if err != nil {
		return
	}
	return
}

func (s ServiceProduct) ListProducts(ctx context.Context, req request.ListProductRequestPayload) (products []entity.Product, err error) {
	queryParam := entity.NewProductQueryFromListProductRequest(req)

	products, err = s.repo.GetAllProduct(ctx, queryParam)

	if err != nil {
		if err == pkg.ErrNotFound {
			return []entity.Product{}, nil
		}
		return
	}

	if len(products) == 0 {
		return []entity.Product{}, nil
	}
	return
}
