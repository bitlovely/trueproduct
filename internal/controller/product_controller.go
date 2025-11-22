package controller

import (
    "net/http"
    "github.com/bitlovely/trueproduct/internal/service"

    "github.com/gin-gonic/gin"
)

// RegisterProduct handles product registration
func RegisterProduct(c *gin.Context) {
    var input struct {
        ProductID string `json:"product_id" binding:"required"`
        Name      string `json:"name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    qrCode, err := service.RegisterProduct(input.ProductID, input.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "service": "trueproduct",
        "message": "Product registered successfully",
        "qr":      qrCode,
    })
}

// VerifyProduct handles QR verification
func VerifyProduct(c *gin.Context) {
    qr := c.Query("qr")
    valid := service.VerifyProduct(qr)

    c.JSON(http.StatusOK, gin.H{
        "service": "trueproduct",
        "valid":   valid,
        "qr":      qr,
    })
}
