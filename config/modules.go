package config

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	//example "github.com/hongyukeji/goravel-modules/example"
)

func init() {
	config := facades.Config()
	config.Add("modules", map[string]any{
		"enabled": config.Env("MODULES_ENABLED", true),

		"providers": []foundation.ServiceProvider{
			//&example.ServiceProvider{},
		},

		"paths": map[string]interface{}{
			"modules":   config.Env("MODULES_PATHS_MODULES", "modules"),
			"assets":    config.Env("MODULES_PATHS_ASSETS", "public"),
			"migration": config.Env("MODULES_PATHS_MIGRATION", "database/migrations"),
		},
	})
}
