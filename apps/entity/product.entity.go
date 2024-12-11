package entity

import (
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id           int       `db:"id"`
	SKU          string    `db:"sku"`
	Name         string    `db:"name"`
	Price        int       `db:"price"`
	Stock        int16     `db:"stock"`
	ID_Categorie int       `db:"id_categorie"`
	Categorie    string    `db:"categorie"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type ProductQuery struct {
	Cursor       int `json:"cursor"`
	Size         int `json:"size"`
	CategoriesID int `json:"categories_id"`
}

func (p Product) Validate() (err error) {
	err = p.ValidateName()
	if err != nil {
		return
	}
	err = p.ValidatePrice()
	if err != nil {
		return
	}
	err = p.ValidateStock()
	if err != nil {
		return
	}
	return
}
func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		err = pkg.ErrProductRequired
		return
	}
	if len(p.Name) < 4 {
		err = pkg.ErrProductInvalid
		return
	}
	return
}
func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		err = pkg.ErrStockInvalid
		return
	}
	return
}
func (p Product) ValidatePrice() (err error) {
	if p.Price <= 0 {
		err = pkg.ErrPriceInvalid
		return
	}
	return
}
func (p Product) ValidateID_Categorie() (err error) {
	if p.ID_Categorie <= 0 {
		err = pkg.ErrIDCategorieInvalid
		return
	}
	return
}

func NewProductFromCreateProductRequest(req request.CreateProductRequestPayload) Product {
	return Product{
		SKU:          uuid.NewString(),
		Name:         req.Name,
		Price:        req.Price,
		Stock:        req.Stock,
		ID_Categorie: req.ID_Categorie,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func NewProductQueryFromListProductRequest(req request.ListProductRequestPayload) ProductQuery {
	req = req.GenerateDefaultValue()
	return ProductQuery{
		Cursor:       req.Cursor,
		Size:         req.Size,
		CategoriesID: req.CategoriesID,
	}
}
