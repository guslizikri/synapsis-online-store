package entity

import (
	"synapsis-online-store/pkg"
	"time"
)

type TransactionStatus uint8

const (
	TransactionStatus_Created    TransactionStatus = 1
	TransactionStatus_Progress   TransactionStatus = 10
	TransactionStatus_InDelivery TransactionStatus = 15
	TransactionStatus_Completed  TransactionStatus = 20

	TRX_CREATED     string = "CREATED"
	TRX_ON_PROGRESS string = "ON_PROGRESS"
	TRX_IN_DELIVERY string = "IN_DELIVERY"
	TRX_COMPLETED   string = "COMPLETED"
	TRX_UNKNOWN     string = "UNKNOWN"
)

var (
	MappingTransactionStatus = map[TransactionStatus]string{
		TransactionStatus_Created:    TRX_CREATED,
		TransactionStatus_Progress:   TRX_ON_PROGRESS,
		TransactionStatus_InDelivery: TRX_IN_DELIVERY,
		TransactionStatus_Completed:  TRX_COMPLETED,
	}
)

type Transactions struct {
	Id           int               `db:"id"`
	UserPublicId string            `db:"user_public_id"`
	TotalPrice   uint              `db:"total_price"`
	Status       TransactionStatus `db:"status"`
	CreatedAt    time.Time         `db:"created_at"`
	UpdatedAt    time.Time         `db:"updated_at"`
}

type TransactionItem struct {
	Id            int       `db:"id"`
	TransactionId uint      `db:"transaction_id"`
	ProductId     uint      `db:"product_id"`
	ProductPrice  uint      `db:"product_price"`
	Quantity      uint8     `db:"Quantity"`
	CreatedAt     time.Time `db:"created_at"`
}

type ProductTrx struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Stock int    `db:"stock" json:"-"`
	Price int    `db:"price" json:"price"`
}

func (t TransactionItem) Validate() (err error) {
	if t.Quantity == 0 {
		return pkg.ErrAmountInvalid
	}
	return
}

func (t TransactionItem) ValidateStock(productStock uint8) (err error) {
	if t.Quantity > productStock {
		return pkg.ErrQuantityGreaterThanStock
	}
	return
}

func (t *Transactions) SetSubTotal(productPrice uint, quantity uint) {
	if t.TotalPrice == 0 {
		t.TotalPrice = productPrice * uint(quantity)
	}
}

func NewTransactionItemFromCreateRequest(productID int, quantity int, price int) TransactionItem {
	return TransactionItem{
		Quantity:     uint8(quantity),
		ProductId:    uint(productID),
		ProductPrice: uint(price),
		CreatedAt:    time.Now(),
	}
}
