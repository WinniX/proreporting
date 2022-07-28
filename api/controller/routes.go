package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/winnix/proreporting/api/config"
	"github.com/winnix/proreporting/api/middleware"
	"github.com/winnix/proreporting/api/templates"
	"github.com/winnix/proreporting/api/web"
)

func Ok(p templates.Page) web.HtmlResult {
	return web.HtmlResult{
		Status: http.StatusOK,
		Page:   p,
	}
}

func BadRequest(json string) web.JsonResult {
	return web.JsonResult{
		Status: http.StatusBadRequest,
		Json:   json,
	}
}

func ServerError() web.HtmlResult {
	return web.HtmlResult{
		Status: http.StatusInternalServerError,
		Page:   &templates.ErrorPage{},
	}
}

func getContext(c *gin.Context) *web.Context {
	ctx := &web.Context{
		Config:    c.MustGet("config").(*config.Config),
		Ctx:       context.Background(),
		GinCtx:    c,
		Validator: c.MustGet("validator").(*validator.Validate),
		Logger:    c.MustGet("logger").(*logrus.Entry),
	}

	ctx.SetCurrentUserIfExists()

	return ctx
}

func wrapper(f web.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)

		res := f(ctx)
		if res == nil {
			return
		}

		res.Render(c)
	}
}

func SetRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("/", wrapper(Index))

	r.GET("/install", wrapper(Install))
	r.GET("/auth/apaleo", wrapper(AuthApaleo))
	r.GET("/auth/apaleo/callback", wrapper(AuthApaleoCallback))

	authorized := r.Group("/")
	authorized.Use(middleware.ApaleoAuthRequired(cfg))
	{
		authorized.GET("/home", wrapper(Home))
	}
}
