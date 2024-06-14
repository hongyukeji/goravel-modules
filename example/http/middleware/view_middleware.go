package middleware

import (
	"html/template"

	gingonic "github.com/gin-gonic/gin"
	"github.com/goravel/framework/contracts/http"
	viewtemplate "github.com/hongyukeji/goravel-modules/example/http/template"
)

// ViewFunc TODO: 开发中
func ViewFunc() http.Middleware {
	return func(ctx http.Context) {
		//facades.View().Share("key", "value")
		//facades.View().Shared("name", facades.Config().GetString("app.name", ""))

		// 方式1
		//context := ctx.(*gin.Context).Instance()
		//facades.View().SetFuncMap(template.FuncMap{"default": viewtemplate.DefaultFunc})

		// 方式2
		router := gingonic.Default()
		// Create a new template and register the custom function
		router.SetFuncMap(template.FuncMap{
			"default": viewtemplate.DefaultFunc,
		})

		ctx.Request().Next()
	}
}
