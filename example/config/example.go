package config

import (
	"github.com/goravel/framework/facades"
	"github.com/hongyukeji/goravel-modules/example/constants"
)

// Boot Start all init methods of the current folder to bootstrap all config.
func Boot() {
	// Boot方法中覆盖app配置有效
	config := facades.Config()
	//config.Add("app.name", config.Env("APP_NAME", "app"))
	config.Add("app.version", config.Env("APP_VERSION", "1.0.0"))
}

func init() {
	// init方法中覆盖app配置无效
	config := facades.Config()
	config.Add("example", map[string]any{
		"name":    config.Env("APP_NAME", constants.Name),
		"version": config.Env("APP_VERSION", constants.Version),
	})
}
