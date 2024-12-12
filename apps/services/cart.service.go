package services

import (
	"context"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
)

type RepoCartIF interface {
	CreateCartItem(ctx context.Context, model entity.CartItemEntity) (err error)
	GetOrCreateCart(ctx context.Context, userPubliID string) (cart entity.Cart, err error)
	GetListCartItem(ctx context.Context, pagination entity.ListCartQuery, cartID int) (Carts []entity.CartItemWithProduct, err error)
	DeleteCartItem(ctx context.Context, cartID int, productID int) (err error)
}

type ServiceCart struct {
	repo RepoCartIF
}

func NewServiceCart(repo RepoCartIF) *ServiceCart {
	return &ServiceCart{
		repo: repo,
	}
}

func (s *ServiceCart) CreateCartItem(ctx context.Context, req request.CreateCartItemRequestPayload) (err error) {
	cart, err := s.repo.GetOrCreateCart(ctx, req.UserPublicID)
	if err != nil {
		return
	}
	cartItemEntity := entity.NewCartItemFromCreateCartItemRequest(req, cart.ID)

	err = cartItemEntity.Validate()
	if err != nil {
		return
	}

	err = s.repo.CreateCartItem(ctx, cartItemEntity)
	if err != nil {
		return
	}

	return
}

func (s *ServiceCart) GetListCartItem(ctx context.Context, req request.PaginationRequestPayload, userPublidID string) (itemProduct []entity.CartItemWithProduct, err error) {
	cart, err := s.repo.GetOrCreateCart(ctx, userPublidID)
	if err != nil {
		return
	}
	pagination := entity.NewCartQueryFromListProductRequest(req)
	itemProduct, err = s.repo.GetListCartItem(ctx, pagination, cart.ID)

	if err != nil {
		if err == pkg.ErrNotFound {
			return []entity.CartItemWithProduct{}, nil
		}
		return
	}

	if len(itemProduct) == 0 {
		return []entity.CartItemWithProduct{}, nil
	}
	return
}

func (s *ServiceCart) DeleteCartItem(ctx context.Context, userPublicID string, productID int) error {
	// Ambil cart user
	cart, err := s.repo.GetOrCreateCart(ctx, userPublicID)
	if err != nil {
		return err
	}

	// Hapus item dari keranjang
	err = s.repo.DeleteCartItem(ctx, cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}
