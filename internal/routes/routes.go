package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/bitlovely/trueproduct/internal/controller"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    api := r.Group("/api/v1/trueproduct")
    {
        api.POST("/register", controller.RegisterProduct)
        api.GET("/verify", controller.VerifyProduct)
    }
}
