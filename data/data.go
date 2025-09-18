// Package data holds the player data which the main CLI will work with
package data

import (
	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

func CreatePlayerData(name string) models.ReportData {
	inventory := []models.Item{
		{Name: "sword", Count: 2},
		{Name: "potion", Count: 3},
		{Name: "bow", Count: 1},
	}

	if !utils.ValidateInventorySlots(inventory) {
		utils.Logger.Error("Default inventory exceeds slot limit", "inventorySize", len(inventory), "maxSlots", utils.MaxInventorySlots)
		if len(inventory) > utils.MaxInventorySlots {
			inventory = inventory[:utils.MaxInventorySlots]
			utils.Logger.Info("Truncated default inventory to fit slot limit", "newSize", len(inventory))
		}
	}

	return models.ReportData{
		Player: models.Player{
			Name:  name,
			Level: 1,
		},
		Inventory: inventory,
	}
}
