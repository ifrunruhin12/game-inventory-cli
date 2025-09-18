// Package data holds the player data which the main CLI will work with
package data

import (
	"bufio"
	"os"
	"strings"

	"github.com/ifrunruhin12/inventory/handlers"
	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

func CreatePlayerData(name string) models.ReportData {
	inventory := []models.Item{
		{Name: "sword", Count: 2},
		{Name: "potion", Count: 3},
		{Name: "bow", Count: 1},
	}

	if !handlers.ValidateInventorySlots(inventory) {
		utils.Logger.Error("Default inventory exceeds slot limit", "inventorySize", len(inventory), "maxSlots", handlers.MaxInventorySlots)
		if len(inventory) > handlers.MaxInventorySlots {
			inventory = inventory[:handlers.MaxInventorySlots]
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

func StartInteractiveSession(tmplMgr *handlers.TemplateManager) error {
	err := tmplMgr.Execute(os.Stdout, "welcome.tmpl", nil)
	if err != nil {
		utils.Logger.Error("Failed to execute welcome template", "error", err)
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		utils.Logger.Error("Failed to read user input", "error", err)
		return err
	}

	name = strings.TrimSpace(name)
	utils.Logger.Info("Player entered name", "name", name)

	playerData := CreatePlayerData(name)

	err = tmplMgr.Execute(os.Stdout, "greeting.tmpl", playerData)
	if err != nil {
		utils.Logger.Error("Failed to execute greeting template", "error", err)
		return err
	}

	commandHandler := handlers.NewCommandHandler(tmplMgr, reader)
	err = commandHandler.HandleCommands(playerData)
	if err != nil {
		utils.Logger.Error("Command handling failed", "error", err)
		return err
	}

	utils.Logger.Info("Interactive session completed successfully")
	return nil
}
