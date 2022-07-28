package controller

import (
	"github.com/winnix/proreporting/api/apaproxy"
	"github.com/winnix/proreporting/api/templates"
	"github.com/winnix/proreporting/api/web"
)

type homeRequest struct {
	AccountCode string `form:"accountCode" validate:"required"`
	SubjectId   string `form:"subjectId" validate:"required"`
}

func Home(ctx *web.Context) web.HandlerResult {
	model := homeRequest{}

	if err := ctx.GinCtx.BindQuery(&model); err != nil {
		ctx.Logger.Infof("failed binding request model: %s\n", err.Error())
		return BadRequest("invalid model")
	}

	if err := ctx.Validator.Struct(model); err != nil {
		ctx.Logger.Infof("invalid request model: %s\n", err.Error())
		return BadRequest("invalid model")
	}

	proxy, err := apaproxy.NewApaProxy(ctx)
	if err != nil {
		ctx.Logger.Errorf("failed initializing apaleo proxy: %s\n", err.Error())
		return ServerError()
	}

	reservations, err := proxy.BookingProxy.GetReservations()
	if err != nil {
		ctx.Logger.Errorf("failed getting reservations: %s\n", err.Error())
		return ServerError()
	}

	ctx.Logger.Infof("number of reservations: %d", len(reservations))

	return Ok(&templates.HomePage{})
}
