package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
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

	return nil
}
