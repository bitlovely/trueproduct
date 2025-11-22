package main

import (
    "github.com/gin-gonic/gin"
    "github.com/bitlovely/trueproduct/internal/routes"
)

func main() {
    r := gin.Default()

    // Load routes
    routes.SetupRoutes(r)

    // Start server
    r.Run(":8080") // http://localhost:8080
}
