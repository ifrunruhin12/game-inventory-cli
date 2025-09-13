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

func CreatePlayerData(name string) models.ReportData {
	return models.ReportData{
		Player: models.Player{
			Name:  name,
			Level: 1,
		},
		Inventory: []models.Item{
			{Name: "sword", Count: 2},
			{Name: "potion", Count: 3},
			{Name: "bow", Count: 1},
		},
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
