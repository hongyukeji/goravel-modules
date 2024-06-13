package modules

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

const Binding = "modules"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})

	if isEnabled() {
		serviceProviders := facades.Config().Get("modules.providers").([]foundation.ServiceProvider)
		for _, serviceProvider := range serviceProviders {
			serviceProvider.Register(app)
		}
	}
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	if isEnabled() {
		serviceProviders := facades.Config().Get("modules.providers").([]foundation.ServiceProvider)
		for _, serviceProvider := range serviceProviders {
			serviceProvider.Boot(app)
		}
	}

	// 将包的配置文件发布到应用程序的 config 目录
	app.Publishes("github.com/hongyukeji/goravel-modules", map[string]string{
		"config/modules.go": app.ConfigPath("modules.go"),
	})
}

func isEnabled() bool {
	return facades.Config().GetBool("modules.enabled")
}
