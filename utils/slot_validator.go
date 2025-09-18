package utils

import "github.com/ifrunruhin12/inventory/models"

const MaxInventorySlots = 5

func ValidateInventorySlots(inventory []models.Item) bool {
	return len(inventory) <= MaxInventorySlots
}

func GetAvailableSlots(inventory []models.Item) int {
	used := len(inventory)
	if used >= MaxInventorySlots {
		return 0
	}
	return MaxInventorySlots - used
}

func IsInventoryFull(inventory []models.Item) bool {
	return len(inventory) >= MaxInventorySlots
}
