package repository

var productStore = make(map[string]string)

// SaveProduct stores a product in memory
func SaveProduct(productID, name string) {
    productStore[productID] = name
}

// ProductExists checks if a product exists
func ProductExists(productID string) bool {
    _, exists := productStore[productID]
    return exists
}
