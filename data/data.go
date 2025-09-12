// Package data holds the player data which it the main CLI will work with
package data

import "github.com/ifrunruhin12/inventory/models"

func DefaultReportData() models.ReportData {
	return models.ReportData{
		Player: models.Player{
			Name:  "Popcycle",
			Level: 7,
		},
		Inventory: []models.Item{
			{Name: "sword", Count: 2},
			{Name: "potion", Count: 3},
			{Name: "bow", Count: 1},
		},
	}
}
