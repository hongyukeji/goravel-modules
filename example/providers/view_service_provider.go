package providers

import (
	"github.com/gin-gonic/gin/render"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	viewtemplate "github.com/hongyukeji/goravel-modules/example/http/template"
)

type ViewServiceProvider struct {
}

func (receiver *ViewServiceProvider) Register(app foundation.Application) {
	// 配置视图模板
	facades.Config().Add(
		"http.drivers.gin.template",
		func() (render.HTMLRender, error) {
			return viewtemplate.DefaultTemplate()
		},
	)
}

func (receiver *ViewServiceProvider) Boot(app foundation.Application) {
	// Add HTTP middleware
	//facades.Route().GlobalMiddleware(middleware.ViewFunc())
}
