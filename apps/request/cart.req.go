package request

type CreateCartItemRequestPayload struct {
	ProductID    int    `db:"product_id" json:"product_id" form:"product_id"`
	UserPublicID string `db:"user_public_id" form:"user_public_id" json:"-"`
	Quantity     int16  `json:"quantity"`
}
