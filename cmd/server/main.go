package main

import (
	"path/filepath"
	"runtime"

	"github.com/bitlovely/trueproduct/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates - get path relative to this file
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	templatesPath := filepath.Join(projectRoot, "templates", "*")
	r.LoadHTMLGlob(templatesPath)

	// Load routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080") // http://localhost:8080
}
