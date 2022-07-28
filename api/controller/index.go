package controller

import (
	"fmt"

	"github.com/winnix/proreporting/api/templates"
	"github.com/winnix/proreporting/api/web"
)

func Index(ctx *web.Context) web.HandlerResult {
	return Ok(&templates.IndexPage{InstallUrl: fmt.Sprintf("%s/install", ctx.Config.Host)})
}
