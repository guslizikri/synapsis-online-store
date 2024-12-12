package repository

import (
	"context"
	"database/sql"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/pkg"

	"github.com/jmoiron/sqlx"
)

type RepoCart struct {
	db *sqlx.DB
}

func NewRepoCart(db *sqlx.DB) *RepoCart {
	return &RepoCart{
		db: db,
	}
}

func (r *RepoCart) CreateCartItem(ctx context.Context, model entity.CartItemEntity) (err error) {
	query := `
			INSERT INTO cart_item (
				product_id, cart_id, quantity, created_at
			) VALUES (
				:product_id, :cart_id, :quantity, :created_at
			)`

	sttm, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer sttm.Close()

	_, err = sttm.ExecContext(ctx, model)
	if err != nil {
		return
	}
	return

}

func (r *RepoCart) GetOrCreateCart(ctx context.Context, userPubliID string) (cart entity.Cart, err error) {
	query := `
		select id, user_public_id, created_at 
		from cart
		where user_public_id = $1
			`
	err = r.db.GetContext(ctx, &cart, query, userPubliID)
	if err == nil {
		return
	}

	// Jika belum ada, buat keranjang baru

	query = `
		INSERT INTO cart (user_public_id) 
		VALUES ($1) 
		RETURNING id`

	sttm, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	sttm.Close()

	// create cart dan return id cart
	// scan dan memasukkannya ke entity cart id
	err = r.db.QueryRowContext(ctx, query, userPubliID).Scan(&cart.ID)
	if err != nil {
		return
	}
	// memasukkan entity userpublicid dengan parameter userpublicid
	cart.UserPublicID = userPubliID

	return

}

func (r *RepoCart) GetListCartItem(ctx context.Context, pagination entity.ListCartQuery, cartID int) (carts []entity.CartItemWithProduct, err error) {
	query := `
		SELECT 
			ci.id, 
			ci.product_id, 
			p.name, 
			p.price,
			ci.quantity
		FROM cart_item ci
		INNER JOIN products p ON ci.product_id = p.id
		WHERE ci.cart_id = $1
		limit $2 offset $3
	`

	err = r.db.SelectContext(ctx, &carts, query, cartID, pagination.Limit, pagination.Offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.ErrNotFound
		}
		return
	}
	return
}

func (r *RepoCart) DeleteCartItem(ctx context.Context, cartID int, productID int) (err error) {
	query := `DELETE FROM cart_item WHERE cart_id = $1 AND product_id = $2`

	_, err = r.db.ExecContext(ctx, query, cartID, productID)
	if err != nil {
		return err
	}

	return nil
}
