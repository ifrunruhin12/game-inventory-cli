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
	for {
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
		case "quit", "exit":
			return ch.handleQuit(playerData)
		case "show":
			err = ch.handleShow(playerData)
			if err != nil {
				return err
			}
		case "cmd":
			err = ch.handleHelp()
			if err != nil {
				return err
			}
		default:
			err = ch.handleUnknownCommand(command)
			if err != nil {
				return err
			}
		}
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
