package constant

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

// method under struct
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// callling temlate  - it will return struct - using pre-defined method with the render
func LoadTemplate() *Template {
	//get app path
	path, _ := os.Executable()
	//get file path
	filePath := filepath.Dir(path)
	//
	templateFolder := fmt.Sprintf("%v/repository/templates/*", filePath)

	template := &Template{
		templates: template.Must(template.ParseGlob(templateFolder)),
	}

	return template
}
