package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
)

type Example struct{}

func NewExample() *Example {
	return &Example{}
}

// Signature The name and signature of the console command.
func (receiver *Example) Signature() string {
	return "example"
}

// Description The console command description.
func (receiver *Example) Description() string {
	return "Example"
}

// Extend The console command extend.
func (receiver *Example) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Example) Handle(ctx console.Context) error {
	fmt.Println("Run Example command")

	var app foundation.Application
	app = facades.App()
	app.Publishes("github.com/hongyukeji/goravel-modules", map[string]string{
		"example/resources/views": app.BasePath("resources/views"),
	}, "example-views")

	return nil
}
