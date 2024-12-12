package services

import (
	"context"
	"synapsis-online-store/apps/repository"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/config"
	"synapsis-online-store/pkg"
	"testing"

	"github.com/stretchr/testify/require"
)

var svcTrx *ServiceTransaction

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
	repo := repository.NewRepoTransaction(db)
	svcTrx = NewServiceTransaction(repo)
}

func TestCreateTransactionSuccess(t *testing.T) {
	var items []request.TransactionItemRequestPayload
	item1 := request.TransactionItemRequestPayload{
		ProductID: 2,
		Quantity:  1,
	}
	item2 := request.TransactionItemRequestPayload{
		ProductID: 1,
		Quantity:  1,
	}
	items = append(items, item1)
	items = append(items, item2)

	req := request.TransactionRequestPayload{
		UserPublicID:    "4ac1501d-8c34-48be-8a7f-e932f07b62fe",
		ItemTransaction: items,
	}

	err := svcTrx.CreateTransaction(context.Background(), req)
	require.Nil(t, err)
}
