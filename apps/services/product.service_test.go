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

var svcProduct *ServiceProduct

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
	repo := repository.NewRepoProduct(db)
	svcProduct = NewServiceProduct(repo)
}

func TestCreateProductSuccess(t *testing.T) {
	req := request.CreateProductRequestPayload{
		Name:  "Indomie",
		Price: 3000,
		Stock: 100,
	}

	err := svcProduct.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}
func TestCreateProductFail(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := request.CreateProductRequestPayload{
			Name:  "",
			Price: 3000,
			Stock: 100,
		}

		err := svcProduct.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, pkg.ErrProductRequired, err)
	})
	t.Run("price must be greater than 0", func(t *testing.T) {
		req := request.CreateProductRequestPayload{
			Name:  "indomie",
			Price: 0,
			Stock: 100,
		}

		err := svcProduct.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, pkg.ErrPriceInvalid, err)
	})
}

func TestListProductSuccess(t *testing.T) {
	pagination := request.ListProductRequestPayload{
		Cursor: 0,
		Size:   10,
	}

	products, err := svcProduct.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}
