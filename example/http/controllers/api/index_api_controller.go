package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

type IndexApiController struct {
	//Dependent services
}

func NewIndexApiController() *IndexApiController {
	return &IndexApiController{
		//Inject services
	}
}

func (r *IndexApiController) Index(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"name":    facades.Config().GetString("app.name", ""),
		"version": facades.Config().GetString("app.version", ""),
		"support": map[string]any{
			"version": support.Version,
		},
	})
}
