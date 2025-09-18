package handlers

import (
	"os"
	"strings"

	"github.com/ifrunruhin12/inventory/models"
	"github.com/ifrunruhin12/inventory/utils"
)

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
		case "slot":
			err = ch.handleSlot(playerData)
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
