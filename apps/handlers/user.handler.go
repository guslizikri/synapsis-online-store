package handlers

import (
	"net/http"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerUser struct {
	svc *services.ServiceUser
}

func NewHandlerUser(svc *services.ServiceUser) *HandlerUser {
	return &HandlerUser{
		svc: svc,
	}
}

func (h *HandlerUser) Register(ctx *fiber.Ctx) (err error) {
	req := request.RegisterRequestPayload{}
	err = ctx.BodyParser(&req)
	if err != nil {
		myErr := pkg.ErrorBadRequest
		return pkg.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			pkg.WithMessage("register fail"),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	err = h.svc.Register(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			pkg.WithMessage("register fail"),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusCreated),
		pkg.WithMessage("register success"),
	).Send(ctx)
}
func (h *HandlerUser) Login(ctx *fiber.Ctx) (err error) {
	req := request.LoginRequestPayload{}
	err = ctx.BodyParser(&req)
	if err != nil {
		myErr := pkg.ErrorBadRequest
		return pkg.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			pkg.WithMessage("login fail"),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	token, err := h.svc.Login(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			pkg.WithMessage("login fail"),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusCreated),
		pkg.WithPayload(map[string]interface{}{
			"access_token": token,
		}),
		pkg.WithMessage("login success"),
	).Send(ctx)
}
