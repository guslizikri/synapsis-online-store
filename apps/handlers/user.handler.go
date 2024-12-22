package handlers

import (
	"log"
	"net/http"
	"strings"
	"synapsis-online-store/apps/request"
	"synapsis-online-store/apps/services"
	"synapsis-online-store/pkg"
	"time"

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

func (h *HandlerUser) Logout(ctx *fiber.Ctx) (err error) {
	authorization := ctx.Get("Authorization")
	if authorization == "" {
		return pkg.NewResponse(
			pkg.WithError(pkg.ErrorUnauthorized),
		).Send(ctx)
	}

	bearer := strings.Split(authorization, "Bearer ")
	if len(bearer) != 2 {
		log.Println("token invalid")
		return pkg.NewResponse(
			pkg.WithError(pkg.ErrorUnauthorized),
		).Send(ctx)
	}

	token := bearer[1]

	expiration := 24 * time.Hour

	// Simpan token ke blacklist melalui service
	err = h.svc.Logout(token, expiration)
	if err != nil {
		myErr, ok := pkg.ErrorMapping[err.Error()]
		if !ok {
			myErr = pkg.ErrorGeneral
		}
		return pkg.NewResponse(
			// withhttpcode akan mengambil dari konfigusrasi witherror
			pkg.WithMessage("logout fail"),
			pkg.WithError(myErr),
		).Send(ctx)
	}

	return pkg.NewResponse(
		pkg.WithHttpCode(http.StatusOK),
		pkg.WithMessage("logout success"),
	).Send(ctx)
}
