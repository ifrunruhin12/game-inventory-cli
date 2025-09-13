// Package handlers handles all command processing and user interaction
package handlers

import (
	"bufio"
	"os"
	"strings"

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

func (ch *CommandHandler) HandleCommands(playerData models.ReportData) error {
	err := ch.tmplMgr.Execute(os.Stdout, "cmd_prompt.tmpl", nil)
	if err != nil {
		utils.Logger.Error("Failed to execute command prompt template", "error", err)
		return err
	}

	command, err := ch.reader.ReadString('\n')
	if err != nil {
		utils.Logger.Error("Failed to read command input", "error", err)
		return err
	}

	command = strings.TrimSpace(command)
	utils.Logger.Info("Player entered command", "command", command)

	switch command {
	case "quit":
		return ch.handleQuit(playerData)
	default:
		return ch.handleUnknownCommand(command)
	}
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
