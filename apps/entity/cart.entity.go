package entity

import (
	"synapsis-online-store/apps/request"
	"synapsis-online-store/pkg"
	"time"
)

type Cart struct {
	ID           int    `db:"id"`
	UserPublicID string `db:"user_public_id"`
	CreatedAt    string `db:"created_at"`
}

type CartItemEntity struct {
	Id        uint      `db:"id"`
	ProductID int       `db:"product_id"`
	CartID    int       `db:"cart_id"`
	Quantity  int16     `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
}

type CartItemWithProduct struct {
	ID          int    `db:"id" json:"id"`
	ProductID   int    `db:"product_id" json:"product_id"`
	ProductName string `db:"name" json:"name"`
	Price       int    `db:"price" json:"price"`
	Quantity    int    `db:"quantity" json:"quantity"`
}

func (c CartItemEntity) Validate() (err error) {
	err = c.ValidateQuantity()
	if err != nil {
		return
	}
	return
}
func (c CartItemEntity) ValidateQuantity() (err error) {
	if c.Quantity <= 0 {
		err = pkg.ErrQuantityInvalid
		return
	}
	return
}

func NewCartItemFromCreateCartItemRequest(req request.CreateCartItemRequestPayload, cart_id int) CartItemEntity {
	return CartItemEntity{
		ProductID: req.ProductID,
		CartID:    cart_id,
		Quantity:  req.Quantity,
		CreatedAt: time.Now(),
	}
}

type ListCartQuery struct {
	Limit  int `json:"limit" query:"limit"`
	Offset int `json:"offset" query:"offset"`
}

func NewCartQueryFromListProductRequest(req request.PaginationRequestPayload) ListCartQuery {
	req = req.GenerateDefaultValue()
	return ListCartQuery{
		Offset: (req.Page - 1) * req.Limit,
		Limit:  req.Limit,
	}
}
