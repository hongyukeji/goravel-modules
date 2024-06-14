package providers

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/hongyukeji/goravel-modules/example/commands"
)

type CommandServiceProvider struct {
}

func (receiver *CommandServiceProvider) Register(app foundation.Application) {
}

func (receiver *CommandServiceProvider) Boot(app foundation.Application) {
	app.Commands([]console.Command{
		commands.NewExample(),
	})
}
