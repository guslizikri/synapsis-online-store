package services

import (
	"context"
	"errors"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
	"time"
)

type RepoTransactionIF interface {
	CreateTransaction(ctx context.Context, trx *entity.Transactions, item []entity.TransactionItem) (err error)
	GetProductByID(ctx context.Context, productID int) (product entity.ProductTrx, err error)
}

type ServiceTransaction struct {
	repo RepoTransactionIF
}

func NewServiceTransaction(repo RepoTransactionIF) *ServiceTransaction {
	return &ServiceTransaction{
		repo: repo,
	}
}

func (s *ServiceTransaction) CreateTransaction(ctx context.Context, req request.TransactionRequestPayload) (err error) {
	var transactionItems []entity.TransactionItem

	var totalPrice int
	for _, item := range req.ItemTransaction {
		product, err := s.repo.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return err
		}
		if product.Id == 0 {
			err = pkg.ErrNotFound
			return err
		}

		if product.Stock < item.Quantity {
			return errors.New("insufficient stock for product " + product.Name)
		}

		// Calculate subtotal for the item
		totalPrice += product.Price * item.Quantity
		trxItem := entity.NewTransactionItemFromCreateRequest(product.Id, item.Quantity, product.Price)

		transactionItems = append(transactionItems, trxItem)
	}

	// Create transaction object
	transaction := entity.Transactions{
		UserPublicId: req.UserPublicID,
		TotalPrice:   uint(totalPrice),
		Status:       entity.TransactionStatus_Created,
		CreatedAt:    time.Now(),
	}

	err = s.repo.CreateTransaction(ctx, &transaction, transactionItems)
	return
}
