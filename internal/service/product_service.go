package service

import (
    "github.com/bitlovely/trueproduct/internal/repository"
    "github.com/bitlovely/trueproduct/internal/utils"
)

// RegisterProduct saves the product and generates a signed QR code
func RegisterProduct(productID, name string) (string, error) {
    // Save product in repository
    repository.SaveProduct(productID, name)

    // Generate signed QR code (HMAC)
    qrCode := utils.GenerateSignedQR(productID)
    return qrCode, nil
}

// VerifyProduct checks if QR code is valid
func VerifyProduct(qr string) bool {
    productID := utils.VerifySignedQR(qr)
    return repository.ProductExists(productID)
}
