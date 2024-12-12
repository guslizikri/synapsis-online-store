package repository

import (
	"context"
	"database/sql"
	"fmt"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/pkg"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	db *sqlx.DB
}

func NewRepoProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{
		db: db,
	}
}

func (r *RepoProduct) CreateProduct(ctx context.Context, model entity.Product) (err error) {
	query := `
		INSERT INTO products (
			sku, name, price, stock, id_categorie, created_at, updated_at
		) Values (
			:sku, :name, :price, :stock, :id_categorie, :created_at, :updated_at
		)
	`
	// stmt ini akan membuka connection pool baru, jadi jangan lupa di close lagi
	stmt, err := r.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model)
	if err != nil {
		return
	}
	return
}

func (r *RepoProduct) GetAllProduct(ctx context.Context, model entity.ProductQuery) (products []entity.Product, err error) {
	filterQuery := ""
	args := []interface{}{model.Cursor} // Placeholder pertama untuk id > $1

	if model.CategoriesID != 0 {
		filterQuery = "AND p.id_categorie = $2"
		args = append(args, model.CategoriesID) // Placeholder kedua untuk categories ID
	}

	// Query utama
	query := fmt.Sprintf(`SELECT 
			p.id, p.sku, p.name,
			p.stock, p.price,
			p.created_at,
			p.updated_at,
			c.categorie
		FROM products p
		JOIN categories c ON c.id = p.id_categorie
		WHERE p.id > $1
		%s
		ORDER BY p.id ASC
		LIMIT $%d`, filterQuery, len(args)+1)

	args = append(args, model.Size)

	// Eksekusi query
	err = r.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.ErrNotFound
		}
		return
	}
	return
}
