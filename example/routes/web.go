package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	controllers "github.com/hongyukeji/goravel-modules/example/http/controllers"
)

func Web() {

	facades.Route().Prefix("example").Group(func(router route.Router) {

		router.Get("", controllers.NewIndexController().Index)

	})

}
