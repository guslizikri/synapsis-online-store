package handlers

import (
	"net/http"
	"strconv"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerCart struct {
	svc *services.ServiceCart
}

func NewHandlerCart(svc *services.ServiceCart) *HandlerCart {
	return &HandlerCart{
		svc: svc,
	}
}

func (h *HandlerCart) CreateCartItem(ctx *fiber.Ctx) error {
	req := request.CreateCartItemRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(ctx)
	}

	userID := ctx.Locals("PUBLIC_ID").(string)
	req.UserPublicID = userID

	err = h.svc.CreateCartItem(ctx.UserContext(), req)
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
		pkg.WithMessage("product added to cart"),
	).Send(ctx)
}
func (h *HandlerCart) GetListCartItem(ctx *fiber.Ctx) error {
	req := request.PaginationRequestPayload{}

	err := ctx.QueryParser(&req)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(ctx)
	}

	userID := ctx.Locals("PUBLIC_ID").(string)

	productItem, err := h.svc.GetListCartItem(ctx.UserContext(), req, userID)
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
		pkg.WithHttpCode(http.StatusOK),
		pkg.WithMessage("get list cart success"),
		pkg.WithPayload(productItem),
		pkg.WithQuery(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h *HandlerCart) DeleteCartItem(ctx *fiber.Ctx) error {
	productIdParam := ctx.Params("product_id", "")
	productID, err := strconv.Atoi(productIdParam)

	publicID := ctx.Locals("PUBLIC_ID").(string)
	err = h.svc.DeleteCartItem(ctx.UserContext(), publicID, productID)
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
		pkg.WithHttpCode(http.StatusOK),
		pkg.WithMessage("product removed from cart"),
	).Send(ctx)
}
