package pkg

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	Query     interface{} `json:"query,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}
	for _, param := range params {
		param(&resp)
	}
	return resp

}

func WithHttpCode(httpcode int) func(r *Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpcode
		return r
	}
}
func WithMessage(message string) func(r *Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}
func WithPayload(payload interface{}) func(r *Response) *Response {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}
func WithQuery(query interface{}) func(r *Response) *Response {
	return func(r *Response) *Response {
		r.Query = query
		return r
	}
}
func WithError(err error) func(r *Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		// type assertion
		myErr, ok := err.(Error)

		if !ok {
			myErr = ErrorGeneral
		}
		r.Error = myErr.Message
		r.ErrorCode = myErr.Code
		r.HttpCode = myErr.HttpCode
		return r
	}
}

func (r Response) Send(ctx *fiber.Ctx) error {
	return ctx.Status(r.HttpCode).JSON(r)
}
