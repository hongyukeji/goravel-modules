package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type IndexController struct {
	//Dependent services
}

func NewIndexController() *IndexController {
	return &IndexController{
		//Inject services
	}
}

func (r *IndexController) Index(ctx http.Context) http.Response {
	return ctx.Response().View().Make("example/index.tmpl", map[string]any{
		"name":    facades.Config().GetString("app.name", ""),
		"version": facades.Config().GetString("app.version", ""),
	})
}
