// Package handlers handles all command processing and user interaction
package handlers

import (
	"bufio"
	"os"

	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

type CommandHandler struct {
	tmplMgr *TemplateManager
	reader  *bufio.Reader
}

func NewCommandHandler(tmplMgr *TemplateManager, reader *bufio.Reader) *CommandHandler {
	return &CommandHandler{
		tmplMgr: tmplMgr,
		reader:  reader,
	}
}

func (ch *CommandHandler) handleShow(playerData models.ReportData) error {
	err := ch.tmplMgr.Execute(os.Stdout, "report.tmpl", playerData)
	if err != nil {
		utils.Logger.Error("Failed to execute report template", "error", err)
		return err
	}
	utils.Logger.Info("Player viewed inventory", "player", playerData.Player.Name)
	return nil
}

func (ch *CommandHandler) handleSlot(playerData models.ReportData) error {
	usedSlots := len(playerData.Inventory)
	availableSlots := utils.GetAvailableSlots(playerData.Inventory)
	isFull := utils.IsInventoryFull(playerData.Inventory)

	slotData := struct {
		UsedSlots      int
		AvailableSlots int
		MaxSlots       int
		IsFull         bool
	}{
		UsedSlots:      usedSlots,
		AvailableSlots: availableSlots,
		MaxSlots:       utils.MaxInventorySlots,
		IsFull:         isFull,
	}

	err := ch.tmplMgr.Execute(os.Stdout, "slot.tmpl", slotData)
	if err != nil {
		utils.Logger.Error("Failed to execute slot template", "error", err)
		return err
	}
	utils.Logger.Info("Player viewed inventory slots", "player", playerData.Player.Name, "usedSlots", usedSlots, "availableSlots", availableSlots, "isFull", isFull)
	return nil
}

func (ch *CommandHandler) handleHelp() error {
	err := ch.tmplMgr.Execute(os.Stdout, "help.tmpl", nil)
	if err != nil {
		utils.Logger.Error("Failed to execute help template", "error", err)
		return err
	}
	utils.Logger.Info("Player viewed available commands")
	return nil
}

func (ch *CommandHandler) handleQuit(playerData models.ReportData) error {
	err := ch.tmplMgr.Execute(os.Stdout, "goodbye.tmpl", playerData)
	if err != nil {
		utils.Logger.Error("Failed to execute goodbye template", "error", err)
		return err
	}
	utils.Logger.Info("Player quit the game", "player", playerData.Player.Name)
	return nil
}

func (ch *CommandHandler) handleUnknownCommand(command string) error {
	commandData := struct {
		Command string
	}{
		Command: command,
	}
	err := ch.tmplMgr.Execute(os.Stdout, "unknown_cmd.tmpl", commandData)
	if err != nil {
		utils.Logger.Error("Failed to execute unknown command template", "error", err)
		return err
	}
	utils.Logger.Info("Player entered unknown command", "command", command)
	return nil
}
