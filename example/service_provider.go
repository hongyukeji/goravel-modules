package example

import (
	"github.com/goravel/framework/contracts/foundation"

	"github.com/hongyukeji/goravel-modules/example/providers"
)

type ServiceProvider struct{}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	(&providers.AppServiceProvider{}).Register(app)
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	(&providers.AppServiceProvider{}).Boot(app)
}
