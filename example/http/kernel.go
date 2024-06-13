package http

import (
	"github.com/goravel/framework/contracts/http"
	//staticmiddleware "github.com/hongyukeji/goravel-static/middleware"
)

type Kernel struct {
}

// The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		//staticmiddleware.Static(),
	}
}
