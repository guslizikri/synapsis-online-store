package pkg

import (
	"errors"
	"net/http"
)

// Error General
var (
	ErrNotFound        = errors.New("data not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
)
var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("email is required")
	ErrPasswordInvalidLength = errors.New("password must have minimun 6 character")
	ErrEmailExist            = errors.New("email already exists")
	ErrPasswordNotMatch      = errors.New("password not match")

	// Product
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product name must have minimum 4 character")
	ErrPriceInvalid    = errors.New("price must be greater than 0")
	ErrStockInvalid    = errors.New("stock must be greater than 0")
	// transactions
	ErrAmountInvalid          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg, code string, httpcode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpcode,
	}
}
func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral         = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "40100", http.StatusForbidden)
)

var (
	// error bad request
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorProductRequired       = NewError(ErrProductRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductInvalid        = NewError(ErrProductInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorStockInvalid          = NewError(ErrStockInvalid.Error(), "40007", http.StatusBadRequest)
	ErrorPriceInvalid          = NewError(ErrPriceInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorInvalidAmount         = NewError(ErrAmountInvalid.Error(), "40009", http.StatusBadRequest)

	ErrorEmailExist       = NewError(ErrEmailExist.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
)

// map ini nantinya untuk pengecekan apakah error yg terjadi,
// ada pada variable error yang sudah dibuat atau tidak
var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrEmailExist.Error():            ErrorEmailExist,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrUnauthorized.Error():          ErrorUnauthorized,
		ErrForbiddenAccess.Error():       ErrorForbiddenAccess,
	}
)
