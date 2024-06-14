package providers

import (
	"github.com/goravel/framework/contracts/foundation"
)

type AppServiceProvider struct {
}

func (receiver *AppServiceProvider) Register(app foundation.Application) {
	(&ConfigServiceProvider{}).Register(app)
	(&AuthServiceProvider{}).Register(app)
	(&RouteServiceProvider{}).Register(app)
	(&ConsoleServiceProvider{}).Register(app)
	(&QueueServiceProvider{}).Register(app)
	(&EventServiceProvider{}).Register(app)
	(&ValidationServiceProvider{}).Register(app)
	(&ViewServiceProvider{}).Register(app)
}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {
	(&ConfigServiceProvider{}).Boot(app)
	(&AuthServiceProvider{}).Boot(app)
	(&RouteServiceProvider{}).Boot(app)
	(&ConsoleServiceProvider{}).Boot(app)
	(&QueueServiceProvider{}).Boot(app)
	(&EventServiceProvider{}).Boot(app)
	(&ValidationServiceProvider{}).Boot(app)
	(&ViewServiceProvider{}).Boot(app)
}
