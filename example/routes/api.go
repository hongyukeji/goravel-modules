package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"

	apis "github.com/hongyukeji/goravel-modules/example/http/controllers/api"
)

func Api() {

	facades.Route().Prefix("api").Group(func(router route.Router) {

		router.Get("/example/ping", func(ctx http.Context) http.Response {
			return ctx.Response().Json(http.StatusOK, map[string]any{
				"name":    facades.Config().GetString("app.name", "app"),
				"version": support.Version,
			})
		})

		router.Get("/example", apis.NewIndexApiController().Index)

	})

}
