package handlers

import (
	"os"
	"strings"

	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

func (ch *CommandHandler) HandleCommands(playerData models.ReportData) error {
	// Use pointer to allow modifications
	data := &playerData

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
		utils.Logger.Info("Player entered command", "command", command, "player", data.Player.Name)

		switch command {
		case "quit", "exit":
			return ch.handleQuit(*data)
		case "show":
			err = ch.handleShow(*data)
			if err != nil {
				return err
			}
		case "slot":
			err = ch.handleSlot(*data)
			if err != nil {
				return err
			}
		case "item":
			err = ch.handleItems()
			if err != nil {
				return err
			}
		case "add":
			err = ch.handleAdd(data)
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
