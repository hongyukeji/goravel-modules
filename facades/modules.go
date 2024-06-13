package facades

import (
	"log"

	"goravel/packages/modules"
	"goravel/packages/modules/contracts"
)

func Modules() contracts.Modules {
	instance, err := modules.App.Make(modules.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Modules)
}
