package handlers

import (
	"net/http"
	"strconv"
	"synapsis-online-store/apps/entity"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerReview struct {
	svc *services.ServiceReview
}

func NewHandlerReview(svc *services.ServiceReview) *HandlerReview {
	return &HandlerReview{
		svc: svc,
	}
}

func (h *HandlerReview) AddReview(c *fiber.Ctx) error {
	var review entity.Review
	err := c.BodyParser(&review)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("invalid payload"),
			pkg.WithError(pkg.ErrorBadRequest),
		).Send(c)
	}

	userID := c.Locals("PUBLIC_ID").(string)
	review.UserPublicID = userID
	err = h.svc.CreateReview(c.Context(), review)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			pkg.WithMessage(err.Error()),
			pkg.WithError(myErr),
		).Send(c)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusCreated),
		pkg.WithMessage("review added successfully"),
	).Send(c)
}

func (h *HandlerReview) GetReviews(c *fiber.Ctx) error {
	productIdParam := c.Params("product_id", "")
	productID, err := strconv.Atoi(productIdParam)
	if err != nil {
		return pkg.NewResponse(
			pkg.WithMessage("failed converted to int"),
			pkg.WithError(pkg.ErrorGeneral),
		).Send(c)
	}

	reviews, err := h.svc.GetReviews(c.Context(), productID)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			pkg.WithMessage(err.Error()),
			pkg.WithError(myErr),
		).Send(c)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusOK),
		pkg.WithMessage("get list review success"),
		pkg.WithPayload(reviews),
	).Send(c)
}
