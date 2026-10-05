package utils

import (
	"rayaw-api/internal/models"

	"github.com/google/uuid"
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
