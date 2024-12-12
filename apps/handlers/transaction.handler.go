package handlers

import (
	"net/http"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerTransaction struct {
	svc *services.ServiceTransaction
}

func NewHandlerTransaction(svc *services.ServiceTransaction) *HandlerTransaction {
	return &HandlerTransaction{
		svc: svc,
	}
}

func (h *HandlerTransaction) CreateTransaction(ctx *fiber.Ctx) error {
	req := request.TransactionRequestPayload{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(ctx)
	}

	userID := ctx.Locals("PUBLIC_ID").(string)
	req.UserPublicID = userID

	err = h.svc.CreateTransaction(ctx.UserContext(), req)
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
		pkg.WithMessage("create transaction success"),
	).Send(ctx)
}
