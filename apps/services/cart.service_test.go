package services

import (
	"context"
	"log"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

var svcCart *ServiceCart

func init() {
	filename := "../../cmd/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := pkg.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		return
	}
	repo := repository.NewRepoCart(db)
	svcCart = NewServiceCart(repo)
}

func TestCreateCartItemSuccess(t *testing.T) {
	req := request.CreateCartItemRequestPayload{
		ProductID:    2,
		UserPublicID: "e54e613b-aa5e-426b-939a-3924cd5e2a4c",
		Quantity:     2,
	}

	err := svcCart.CreateCartItem(context.Background(), req)
	require.Nil(t, err)
}

func TestListCartItemSuccess(t *testing.T) {
	pagination := request.PaginationRequestPayload{
		Page:  0,
		Limit: 10,
	}
	// pastikan datanya ada terlebih dahulu di db
	UserPublicID := "e54e613b-aa5e-426b-939a-3924cd5e2a4c"

	cartItemProducts, err := svcCart.GetListCartItem(context.Background(), pagination, UserPublicID)
	require.Nil(t, err)
	require.NotNil(t, cartItemProducts)
	log.Printf("%+v", cartItemProducts)
}
func TestDeleteCartItemSuccess(t *testing.T) {
	// pastikan datanya ada terlebih dahulu di db
	userPublicID := "e54e613b-aa5e-426b-939a-3924cd5e2a4c"
	productID := 2

	err := svcCart.DeleteCartItem(context.Background(), userPublicID, productID)
	require.Nil(t, err)
}
