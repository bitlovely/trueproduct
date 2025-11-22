package utils

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
)

// Secret key for HMAC (use env variable in production)
var secretKey = []byte("supersecretkey")

// GenerateSignedQR returns a base64 HMAC of the productID
func GenerateSignedQR(productID string) string {
    mac := hmac.New(sha256.New, secretKey)
    mac.Write([]byte(productID))
    signature := mac.Sum(nil)
    return base64.StdEncoding.EncodeToString([]byte(productID + ":" + string(signature)))
}

// VerifySignedQR returns productID if signature is valid, else empty string
func VerifySignedQR(qr string) string {
    data, err := base64.StdEncoding.DecodeString(qr)
    if err != nil {
        return ""
    }
    parts := string(data)
    for i := 0; i < len(parts); i++ {
        if parts[i] == ':' {
            productID := parts[:i]
            sig := parts[i+1:]
            mac := hmac.New(sha256.New, secretKey)
            mac.Write([]byte(productID))
            expectedSig := string(mac.Sum(nil))
            if hmac.Equal([]byte(sig), []byte(expectedSig)) {
                return productID
            }
            break
        }
    }
    return ""
}
