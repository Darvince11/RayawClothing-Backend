package utils

import (
	"rayaw-api/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func MergeMap(m1, m2 *map[string]any) map[string]any {
	//This function appends m2 to m1
	result := make(map[string]any)
	for k, v := range *m1 {
		result[k] = v
	}
	for k, v := range *m2 {
		result[k] = v
	}
	return result
}

func GroupItemsById(items *[]models.OrderItem) map[uuid.UUID][]models.OrderItem {
	//This function takes *[]models.OrderItem and groups them by orderId and returns a map of orderId to []models.OrderItem
	result := make(map[uuid.UUID][]models.OrderItem)
	for _, item := range *items {
		result[item.OrderId] = append(result[item.OrderId], item)
	}
	return result
}

func Hash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)

	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func CheckHash(hash string, text string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
}
