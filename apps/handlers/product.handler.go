package handlers

import (
	"net/http"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/apps/response"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerProduct struct {
	svc *services.ServiceProduct
}

func NewHandlerProduct(svc *services.ServiceProduct) *HandlerProduct {
	return &HandlerProduct{
		svc: svc,
	}
}

func (h *HandlerProduct) CreateProduct(ctx *fiber.Ctx) error {
	req := request.CreateProductRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(ctx)
	}

	err = h.svc.CreateProduct(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			pkg.WithMessage(err.Error()),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusCreated),
		pkg.WithMessage("create product success"),
	).Send(ctx)
}
func (h *HandlerProduct) GetListProduct(ctx *fiber.Ctx) error {
	req := request.ListProductRequestPayload{}

	err := ctx.QueryParser(&req)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(ctx)
	}

	products, err := h.svc.ListProducts(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			pkg.WithMessage(err.Error()),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	productListResponse := response.NewProductListResponseFromEntity(products)

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusOK),
		pkg.WithMessage("get list product success"),
		pkg.WithPayload(productListResponse),
		pkg.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}
