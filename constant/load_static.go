package constant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	//load static path
	path, _ := os.Executable()
	//get file path
	filepath := filepath.Dir(path)
	//
	staticFolder := fmt.Sprintf("%v/repository/assets", filepath)
	app.Static("static", staticFolder)
}
