package providers

import (
	"github.com/goravel/framework/contracts/foundation"

	"github.com/hongyukeji/goravel-modules/example/config"
)

type ConfigServiceProvider struct {
}

func (receiver *ConfigServiceProvider) Register(app foundation.Application) {
	config.Boot()
}

func (receiver *ConfigServiceProvider) Boot(app foundation.Application) {
	//config.Boot()
}
