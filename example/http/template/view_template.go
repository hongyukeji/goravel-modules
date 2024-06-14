package template

import (
	"html/template"

	"github.com/gin-gonic/gin/render"
	"github.com/goravel/gin"
)

// DefaultTemplate creates a TemplateRender instance with default options.
func DefaultTemplate() (*render.HTMLProduction, error) {
	renderOptions := gin.RenderOptions{
		FuncMap: template.FuncMap{
			"default": DefaultFunc,
		},
	}
	return gin.NewTemplate(renderOptions)
}

// DefaultFunc returns the first non-empty string from the arguments
func DefaultFunc(defaultValue interface{}, value interface{}) interface{} {
	if value == nil || value == "" {
		return defaultValue
	}
	return value
}
