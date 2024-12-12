package repository

import (
	"context"
	"database/sql"
	"fmt"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/pkg"

	"github.com/jmoiron/sqlx"
)

type RepoTransaction struct {
	db *sqlx.DB
}

func NewRepoTransaction(db *sqlx.DB) *RepoTransaction {
	return &RepoTransaction{
		db: db,
	}
}
func (r *RepoTransaction) CreateTransaction(ctx context.Context, transaction *entity.Transactions, items []entity.TransactionItem) error {
	// Start transaction
	tx, err := r.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Insert into transactions table
	var transactionID int64
	err = tx.QueryRowContext(ctx,
		`INSERT INTO transactions (user_public_id, total_price, status) 
		   VALUES ($1, $2, $3) RETURNING id`,
		transaction.UserPublicId, transaction.TotalPrice, transaction.Status,
	).Scan(&transactionID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert into transaction_item table and update stock
	for _, item := range items {
		// Insert transaction item
		_, err = tx.ExecContext(ctx,
			`INSERT INTO transaction_item (transaction_id, product_id, quantity, product_price) 
			   VALUES ($1, $2, $3, $4)`,
			transactionID, item.ProductId, item.Quantity, item.ProductPrice,
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return fmt.Errorf("rollback error: %v, original error: %v", rollbackErr, err)
			}
			return err
		}

		// Update product stock
		_, err = tx.ExecContext(ctx,
			`UPDATE products
			   SET stock = stock - $1
			   WHERE id = $2 AND stock >= $1`,
			item.Quantity, item.ProductId,
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return fmt.Errorf("rollback error: %v, original error: %v", rollbackErr, err)
			}
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}

func (r *RepoTransaction) GetProductByID(ctx context.Context, productID int) (product entity.ProductTrx, err error) {
	query := `
		SELECT 
			id, name, stock, price
		FROM products
		WHERE id=$1
	`

	err = r.db.GetContext(ctx, &product, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.ProductTrx{}, pkg.ErrNotFound
		}
		return
	}

	return
}
