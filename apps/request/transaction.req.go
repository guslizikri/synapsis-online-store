package request

type TransactionRequestPayload struct {
	UserPublicID    string                          `json:"-"`
	ItemTransaction []TransactionItemRequestPayload `json:"item_transaction"`
}

type TransactionItemRequestPayload struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
